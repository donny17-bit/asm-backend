package claim

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


func Accepted(c *gin.Context)  {
	session := sessions.Default(c)
	ldc_id := session.Get("ldc_id")

	fmt.Println("accepted controller")

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
	order := "KANWIL, CABANG, PERWAKILAN, [SUB PERWAKILAN], NAMALEADER0, NAMALEADER1, NAMALEADER2, NAMALEADER3"  // order default (req)
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
	queryRow := "SELECT COUNT(1) FROM ACCEPTEDCLAIM_GABUNGAN_VIEW A INNER JOIN MV_AGEN MA ON A.LAG_AGEN_ID = MA.LAG_AGEN_ID INNER JOIN LST_CABANG LC ON A.LDC_ID = LC.LDC_ID INNER JOIN LST_BUSINESS LB ON A.LBU_ID = LB.LBU_ID INNER JOIN LST_GRP_BUSINESS LGB ON LB.LGB_ID = LGB.LGB_ID LEFT OUTER JOIN Warehouse_Asm.dbo.MST_CLIENT MC ON A.CLIENT_ID = MC.MCL_ID "
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

		whereDate := " and CAST(left(ACCEPTED_DATE, 4) + right(ACCEPTED_DATE, 2) + left(right(ACCEPTED_DATE, 5), 2)  AS INT) between  '" + formatedBeginDate + "' and '" + formatedEndDate + "' "
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
	queryFinal := "exec SP_DETAIL_ACCEPTED_CLAIM " + " "
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
		whereDate := " and CAST(left(ACCEPTED_DATE, 4) + right(ACCEPTED_DATE, 2) + left(right(ACCEPTED_DATE, 5), 2)  AS INT) between ''" + formatedBeginDate + "'' and ''" + formatedEndDate + "'' "
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

		// Scan each row into a struct
		err := rows.Scan(
			&data.Rn,
			&data.AcceptedDateOri,
			&data.Kanwil,
			&data.Cabang,
			&data.Perwakilan,
			&data.SubPerwakilan,
			&data.Namaleader0,
			&data.Namaleader1,
			&data.Namaleader2,
			&data.Namaleader3,
			&data.GroupBusiness,
			&data.Business,
			&data.TahunPolis,
			&data.AcceptedNo,
			&data.NoKlaim,
			&data.NoPolis,
			&data.NoCif,
			&data.ClientName,
			&data.Mo,
			&data.PrepareDate,
			&data.DateOfLoss,
			&data.AcceptedDate,
			&data.BeginDate,
			&data.EndDate,
			&data.JenisPaket,
			&data.Workshop,
			&data.NamaDealer,
			&data.ColDesk,
			&data.RiskLoc,
			&data.Tsi,
			&data.AcceptedClaim,
			&data.AcceptedClaimRp,
			&data.AccKlaimGrossRp,
			&data.OwnRetention,
			&data.CoIns,
			&data.Psrspl,
			&data.Qsri,
			&data.Er1,
			&data.Surplus1,
			&data.Surplus2,
			&data.Er2,
			&data.PsrqsRi,
			&data.PsrqsOr,
			&data.Ors,
			&data.Facultative,
			&data.Facobl,
			&data.Bppdan,
			&data.Xl,
			&data.Pss,
			&data.Prgbi,
			&data.Pfra,
			&data.Fsplnsri,
			&data.Psplnsri,
			&data.Fsplnsor,
			&data.Psplnsor,
			&data.Facobsrb,
			&data.Facobindt,
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