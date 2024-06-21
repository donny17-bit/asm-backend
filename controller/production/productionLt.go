package production

import (
	"asm-backend/model"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func ProductionLt(c *gin.Context) {
	session := sessions.Default(c)
	ldc_id := session.Get("ldc_id") // sesuai login cabang user

	id := session.Get("id")
	nik := session.Get("nik")
	lastActivity := session.Get("lastActivity")
	expiration := session.Get("expiration")

	fmt.Println("production controller")
	fmt.Println("session : ", session)
	fmt.Println("id : ", id)
	fmt.Println("nik : ", nik)
	fmt.Println("ldc_id : ", ldc_id)
	fmt.Println("lastActivity : ", lastActivity)
	fmt.Println("expiration : ", expiration)


	if ldc_id == nil {
		fmt.Println("error cabang kosong")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        400,
			"data":          nil,
			"message":       "Failed get data, cabang is not present in session",
		})
		return
	}

	var inputData InputData
	if err := c.BindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page := inputData.Page          // req
	pageSize := inputData.Page_size // req
	sort := ""                   // opt
	order := "thnbln, client_name"  // order default (req)
	noPolis := inputData.No_polis
	noCif := inputData.No_cif
	beginDate := inputData.Begin_date
	endDate := inputData.End_date
	business := inputData.Business
	businessSource := inputData.Sumbis
	clientName := inputData.Client_name
	branch := inputData.Branch

	var formatedBeginDate string
	var formatedEndDate string

	fmt.Println("ldc_id", ldc_id)
	fmt.Println("noPolis", noPolis)
	fmt.Println("noCif", noCif)
	fmt.Println("beginDate", beginDate)
	fmt.Println("endDate", endDate)
	fmt.Println("business", business)
	fmt.Println("clientName", clientName)

	if sort == "" {
		sort = "asc"
	}

	db, err := model.SqlModel()

	if err != nil {
		fmt.Println("error to connect to database")
		return
	}
	defer db.Close()

	// field begin date & end date harus di isi
	if beginDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{ // nnti status ok nya di ganti status failed
			"status":  400,
			"data":    "",
			"message": "Failed get data, please provide valid date periode",
		})
		return
	}

	// cek total row ----------------------------------------------------------------
	queryRow := "select count(1) as total_rows FROM PRODUCTION_GABUNGAN_VIEW A JOIN MV_AGEN MA ON A.LAG_ID = MA.LAG_AGEN_ID JOIN LST_CABANG LC ON A.LDC_ID = LC.LDC_ID JOIN LST_BUSINESS LB ON A.LBU_ID = LB.LBU_ID JOIN LST_GRP_BUSINESS LGB ON LB.LGB_ID = LGB.LGB_ID JOIN LST_JN_PRODUKSI LJP ON LJP.LJP_ID = A.LJP_ID JOIN JNNER JNN ON JNN.JN_NER = A.JN_NER LEFT OUTER JOIN LST_MO MO ON A.LMO_ID = MO.LMO_ID LEFT OUTER JOIN MST_CLIENT MC ON A.CLIENT_ID = MC.MCL_ID LEFT OUTER JOIN LST_JENIS_COAS JN_COAS ON A.MDS_JN_COAS = JN_COAS.MDS_JN_COAS "
	whereRow := "where a.ldc_id = '" + ldc_id.(string) + "' "

	queryRow = queryRow + whereRow

	// filter polis
	if noPolis != "" {
		andPolis := " and no_polis in ('" + noPolis + "','') "
		queryRow = queryRow + andPolis
	}

	// filter no cif
	if noCif != "" {
		andCif := " and no_cif in ('" + noCif + "','') "
		queryRow = queryRow + andCif
	}

	// filter bisnis
	if business != "" {
		whereBusiness := " and (LBU_NOTE like '%" + business + "%' OR LGB_NOTE like '%" + business + "%') "
		queryRow = queryRow + whereBusiness
	}

	// filter client name
	if clientName != "" {
		whereClient := " and MC.MCL_NAME like '%" + clientName + "%' "
		queryRow = queryRow + whereClient
	}

	// filter sumbis
	if businessSource != "" {
		whereBusinessSource := " and (MA.NAMALEADER0 like '%" + businessSource + "%' OR " +
			"MA.NAMALEADER1 like '%" + businessSource + "%' OR " +
			"MA.NAMALEADER2 like '%" + businessSource + "%' OR " + 
			"MA.NAMALEADER3 like '%" + businessSource + "%') "
		queryRow = queryRow + whereBusinessSource
	}

	// filter branch
	if branch != "" {
		whereBranch := " and (LC.KANWIL LIKE '%" + branch + "%' OR " + 
			"LC.CABANG LIKE '%" + branch + "%' OR "+
			"LC.PERWAKILAN LIKE '%" + branch + "%' OR "+
			"LC.SUB_PERWAKILAN LIKE '%" + branch + "%') "
		queryRow = queryRow + whereBranch
	}

	// filter tgl
	if beginDate != "" && endDate != "" {
		// change date format
		parsedBeginDate, err := time.Parse("2006-01-02", beginDate)
		parsedEndDate, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}

		formatedBeginDate = parsedBeginDate.Format("20060102")
		formatedEndDate = parsedEndDate.Format("20060102")

		whereDate := " and CAST(left(tgl_prod, 4) + right(TGL_PROD, 2) + left(right(TGL_PROD, 5), 2)  AS INT) between  '" + formatedBeginDate + "' and '" + formatedEndDate + "' "
		queryRow = queryRow + whereDate
	}

	fmt.Println("queryRow row : ", queryRow)

	countRows, err := db.Query(queryRow)

	if err != nil {
		fmt.Println("Error executing query row:", err)
		return
	}

	defer countRows.Close()

	var totalRows int

	// save query result to total rows variable
	for countRows.Next() {
		err := countRows.Scan(&totalRows)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{ 
				"status":  400,
				"data":    "",
				"message": "Failed get data, eror on the server "+err.Error(),
			})
			return
		}
	}

	fmt.Println("total rows", totalRows)

	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil {
		fmt.Println("Error convert page size to int 1 :", err)
		return
	}

	totalPage := math.Ceil(float64(totalRows) / float64(pageSizeNum))
	// // totalPageFloat := float64(totalPage)

	if err := countRows.Err(); err != nil {
		fmt.Println("Error iterating count rows:", err)
		return
	}

	// end count rows
	// ------------------------------------------------------------------------------------------------------------

	// get query
	// var queryFinal string
	queryFinal := "exec SP_DETAIL_PRODUCTION_LONGTERM " + " "
	where := "'" + order + "', '" + sort + "', '" + page + "', '" + pageSize + "', 'where a.ldc_id = ''" + ldc_id.(string) + "''" + " "

	queryFinal = queryFinal + where

	// filter polis
	if noPolis != "" {
		andPolis := " and no_polis in (''" + noPolis + "'','''')"
		queryFinal = queryFinal + andPolis
	}

	// filter no cif
	if noCif != "" {
		andCif := " and no_cif in (''" + noCif + "'','''')"
		queryFinal = queryFinal + andCif
	}

	// filter bisnis
	if business != "" {
		whereBusiness := " and (LBU_NOTE like ''%" + business + "%'' OR LGB_NOTE like ''%" + business + "%'') "
		queryFinal = queryFinal + whereBusiness
	}

	// filter client name
	if clientName != "" {
		whereClient := " and MCL_NAME like ''%" + clientName + "%'' "
		queryFinal = queryFinal + whereClient
	}

	// filter sumbis
	if businessSource != "" {
		whereBusinessSource := " and (MA.NAMALEADER0 like ''%" + businessSource + "%'' OR " +
			"MA.NAMALEADER1 like ''%" + businessSource + "%'' OR " +
			"MA.NAMALEADER2 like ''%" + businessSource + "%'' OR " + 
			"MA.NAMALEADER3 like ''%" + businessSource + "%'') "
		queryFinal = queryFinal + whereBusinessSource
	}

	// filter branch
	if branch != "" {
		whereBranch := " and (LC.KANWIL LIKE ''%" + branch + "%'' OR " + 
			"LC.CABANG LIKE ''%" + branch + "%'' OR "+
			"LC.PERWAKILAN LIKE ''%" + branch + "%'' OR "+
			"LC.SUB_PERWAKILAN LIKE ''%" + branch + "%'') "
		queryFinal = queryFinal + whereBranch
	}


	// filter tgl
	if beginDate != "" && endDate != "" {
		whereDate := " and CAST(left(tgl_prod, 4) + right(TGL_PROD, 2) + left(right(TGL_PROD, 5), 2)  AS INT) between ''" + formatedBeginDate + "'' and ''" + formatedEndDate + "'' "
		queryFinal = queryFinal + whereDate
	}

	queryFinal = queryFinal + "'"

	fmt.Println("queryFinal get : ", queryFinal)
	// nnti tambhain klau yg login nik itasm, keluarin semua
	rows, err := db.Query(queryFinal)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close() // Close the result set when done


	// Create an array to store the query results
	var datas []Data

	for rows.Next() {
		var data Data
		// var JenisPaketSql sql.NullString
		// var NamaCedingSql sql.NullString

		// Scan each row into a struct
		err := rows.Scan(&data.Rn,
			&data.TglProd,
			&data.ThnBln,
			&data.ProdDate,
			&data.BeginDate,
			&data.EndDate,
			&data.Mo,
			&data.ClientName,
			&data.Kanwil,
			&data.Cabang,
			&data.Perwakilan,
			&data.SubPerwakilan,
			&data.Jnner,
			&data.JenisProd,
			// &JenisPaketSql,
			&data.JenisPaket,
			&data.Ket,
			// &NamaCedingSql,
			&data.NamaCeding,
			&data.Namaleader0,
			&data.Namaleader1,
			&data.Namaleader2,
			&data.Namaleader3,
			&data.GroupBusiness,
			&data.Business,
			&data.NoKontrak,
			// &data.Alasan,
			&data.NoPolis,
			&data.NoCif,
			&data.ProdKe,
			&data.NamaDealer,
			&data.Tsi,
			&data.Gpw,
			&data.Disc,
			&data.Disc2,
			&data.Comm,
			&data.Oc,
			&data.Bkp,
			&data.Ngpw,
			&data.Ri,
			&data.Ricom,
			&data.Npw,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{ 
				"status":  400,
				"data":    "",
				"message": "Failed get data, eror on the server "+err.Error(),
			})
			return
		}

		// Append the struct to the array
		datas = append(datas, data)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}

	// --------- end tarik data from sql ----------

	pageNum, err := strconv.Atoi(page)
	// if err != nil {
	// 	// Handle error
	// 	fmt.Println("Error converting page to number:", err)
	// 	return
	// }
	nextPage := pageNum + 1

	// check if next page is greater than total page or not
	if nextPage > int(totalPage) {
		nextPage = int(totalPage)
	}

	previousPage := pageNum - 1

	currentPage, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println("Error convert page size to int 2 :", err)
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		fmt.Println("Error convert page size to int 3 :", err)
		return
	}

	// Respond with JSON data
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"data":          datas,
		"current_page":  currentPage,
		"next_page":     nextPage,
		"previous_page": previousPage,
		"max_page":      totalPage,
		"page_size":     pageSizeInt,
		"total_data":    totalRows,
		"message":       "success get data",
	})

}
