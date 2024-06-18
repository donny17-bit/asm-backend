package production

import (
	"asm-backend/controller/auth"
	"asm-backend/model"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
)


func ExportProdYr(c *gin.Context) {
	err := godotenv.Load()

	if err != nil {
		fmt.Print("Load env failed")
		return
	}

	auth_server := os.Getenv("auth_server")

	var ok bool

	if auth_server == "oracle" {
		ok = auth.IsActive(c)
	}

	if auth_server == "sql" {
		ok = auth.IsActiveSql(c)
	}

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"data":           "",
			"current_page: ": "",
			"page_size":      "",
			"max_page":       "",
			"message":        "unauthorized",
			"status":         401,
		})
		return
	}

	session := sessions.Default(c)
	ldc_id := session.Get("ldc_id") // default sesuai info login

	if ldc_id == nil {
		fmt.Println("error cabang kosong")
		return
	}

	var inputData InputData
	if err := c.BindJSON(&inputData); err != nil {
		fmt.Println("error : ", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek jika ldc_id ada di request
	var ldc_id_param string

	if c.Query("ldc_id") == "" {
		ldc_id_param = ldc_id.(string)
	} else {
		ldc_id_param = c.PostForm("ldc_id") // nnti diganti
	}

	page := "0"                    // req
	pageSize := "0"                // req
	sort := "asc"                  // opt
	order := "thnbln, client_name" // req
	noPolis := inputData.No_polis
	noCif := inputData.No_cif
	beginDate := inputData.Begin_date
	endDate := inputData.End_date
	business := inputData.Business
	businessSource := inputData.Sumbis
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

	// field begin & end date harus di isi
	if beginDate == "" || endDate == "" {
		c.JSON(http.StatusOK, gin.H{ // nnti status ok nya di ganti status failed
			"status":  400,
			"data":    "",
			"message": "Failed get data, please provide valid date periode",
		})
		return
	}

	// get query
	// var queryFinal string
	queryFinal := "exec SP_DETAIL_PRODUCTION_YEAR " + " "
	where := "'" + order + "', '" + sort + "', '" + page + "', '" + pageSize + "', 'where a.ldc_id = ''" + ldc_id_param + "''" + " "

	queryFinal = queryFinal + where

	// filter polis
	if noPolis != "" {
		andPolis := " and no_polis in (''" + noPolis + "'','''')"
		queryFinal = queryFinal + andPolis
	}

	// filter polis
	if noCif != "" {
		andCif := " and no_cif in (''" + noCif + "'','''')"
		queryFinal = queryFinal + andCif
	}

	// filter bisnis
	if business != "" {
		whereBusiness := " and (LBU_NOTE like ''%" + business + "%'' OR LGB_NOTE like ''%" + business + "%'') "
		queryFinal = queryFinal + whereBusiness
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

		fmt.Println("begin_date : ", beginDate)
		fmt.Println("end_date : ", endDate)

		// change date format
		parsedBeginDate, err := time.Parse("2006-01-02", beginDate)
		parsedEndDate, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}

		formatedBeginDate := parsedBeginDate.Format("20060102")
		formatedEndDate := parsedEndDate.Format("20060102")

		whereDate := " and CAST(left(tgl_prod, 4) + right(TGL_PROD, 2) + left(right(TGL_PROD, 5), 2)  AS INT) between ''" + formatedBeginDate + "'' and ''" + formatedEndDate + "'' "
		queryFinal = queryFinal + whereDate
	}

	queryFinal = queryFinal + "'"
	fmt.Println("queryFinal : ", queryFinal)

	// Execute the query
	rows, err := db.Query(queryFinal)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// Create a new Excel file
	file := excelize.NewFile()
	sheetName := "Sheet1"
	// index := file.NewSheet(sheetName)

	// Fetch the column names
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}

	// Write column headers to the Excel file
	for colIndex, colName := range columns {
		cell := toAlphaString(colIndex+1) + "1"
		file.SetCellValue(sheetName, cell, colName)
	}

	// Write data rows to the Excel file
	rowIndex := 2
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			fmt.Println(err)
		}
		for colIndex, value := range values {
			cell := toAlphaString(colIndex+1) + strconv.Itoa(rowIndex)
			file.SetCellValue(sheetName, cell, value)
		}
		rowIndex++
	}

	// Save the Excel file
	tempFile := "Detail_Produksi_Yearly.xlsx" // Temporarily save the file
	if err := file.SaveAs(tempFile); err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Serve the Excel file as a response
	defer os.Remove(tempFile) // Remove the file after serving
	c.Header("Content-Disposition", "attachment; filename=Detail_Produksi_Yearly.xlsx")
	c.Header("Content-Type", "application/octet-stream")
	c.File(tempFile)
}
