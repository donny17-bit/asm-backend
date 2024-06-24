package surplus

import (
	"asm-backend/model"
	"fmt"
	"math"
	"net/http"
	"strconv"

	// "time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SurplusYr(c *gin.Context) {
	session := sessions.Default(c)
	ldc_id := session.Get("ldc_id") // sesuai login cabang user

	fmt.Println("surplus yearly controller")
	fmt.Println("ldc_id : ", ldc_id)

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
	order := "kanwil, cabang"  // order default (req)
	noPolis := inputData.No_polis
	noCif := inputData.No_cif
	periode := inputData.Periode
	business := inputData.Business
	businessSource := inputData.Sumbis
	clientName := inputData.Client_name
	branch := inputData.Branch

	if sort == "" {
		sort = "asc"
	}

	db, err := model.SqlModel()

	if err != nil {
		fmt.Println("error to connect to database")
		return
	}
	defer db.Close()

	// cek total row ----------------------------------------------------------------
	queryRow := "select count(1) as total_row FROM prodclaimuwmththn_gabungan_view A INNER JOIN MV_AGEN MA ON A.LAG_ID = MA.LAG_AGEN_ID INNER JOIN LST_CABANG LC ON A.LDC_ID = LC.LDC_ID INNER JOIN LST_BUSINESS LB ON A.LBU_ID = LB.LBU_ID INNER JOIN LST_GRP_BUSINESS LGB ON LB.LGB_ID = LGB.LGB_ID LEFT OUTER JOIN LST_MO MO ON A.LMO_ID = MO.LMO_ID LEFT OUTER JOIN mst_client mc ON A.client_ID = mc.mcl_ID LEFT OUTER JOIN Warehouse_Asm.dbo.LST_JENIS_COAS JN_COAS ON A.MDS_JN_COAS = JN_COAS.MDS_JN_COAS "
	whereRow := "where a.ldc_id = '" + ldc_id.(string) + "' and left(mthname, len(mthname) - 3) = '"+periode+"' "
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

	if err := countRows.Err(); err != nil {
		fmt.Println("Error iterating count rows:", err)
		return
	}

	// end count row
	// ------------------------------------


	// get query
	// var queryFinal string
	queryFinal := "exec SP_SURPLUS_LONGTERM_CIF " + " "
	where := " '"+ order +"', '"+ sort +"', '"+ page +"', '"+ pageSize +"', 'where a.ldc_id = ''"+ ldc_id.(string) +"'' and left(mthname, len(mthname) - 3) = ''"+ periode +"''  "

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

	queryFinal = queryFinal + "'"
	fmt.Println("queryFinal get : ", queryFinal)

	rows, err := db.Query(queryFinal)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close()

	// Create an array to store the query results
	var datas []Data

	for rows.Next() {
		var data Data
		// var JenisPaketSql sql.NullString
		// var NamaCedingSql sql.NullString

		// Scan each row into a struct
		err := rows.Scan(
			&data.Rn,
			&data.Mthname,
			&data.Periode,
			&data.Kanwil,
			&data.Cabang,
			&data.Perwakilan,
			&data.SubPerwakilan,
			&data.Namaleader0,
			&data.Namaleader1,
			&data.Namaleader2,
			&data.Namaleader3,
			&data.Mo,
			&data.GroupBusiness,
			&data.Business,
			&data.ClientName,
			&data.NoPolis,
			&data.NoCif,
			&data.JenisPaket,
			&data.Keterangan,
			&data.NamaCeding,
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
			&data.CadPremi,
			&data.CadPremi1,
			&data.PremiumReserve,
			&data.Npe,
			&data.AcceptedClaim,
			&data.RejectedClaim,
			&data.OutstandingClaim,
			&data.ReversedClaim,
			&data.SurplusUw,
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