package web

import (
	"asm-backend/controller/auth"
	"asm-backend/model"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
)

// Convert column number to Excel column letter
func toAlphaString(n int) string {
	if n <= 0 {
		return ""
	}
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaLen := len(alpha)
	result := ""
	for n > 0 {
		index := (n - 1) % alphaLen
		result = string(alpha[index]) + result
		n = (n - 1) / alphaLen
	}
	return result
}

func ExportProdLt(c *gin.Context) {
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

	// cek jika ldc_id ada di request
	var ldc_id_param string

	if c.Query("ldc_id") == "" {
		ldc_id_param = ldc_id.(string)
	} else {
		ldc_id_param = c.Query("ldc_id")
	}

	page := c.Query("page")          // req
	pageSize := c.Query("page_size") // req
	sort := c.Query("sort")          // opt
	order := c.Query("order")        // req
	noPolis := c.Query("no_polis")
	beginDate := c.Query("begin_date")
	endDate := c.Query("end_date")
	business := c.Query("business")

	if sort == "" {
		sort = "asc"
	}

	db, err := model.SqlModel()

	if err != nil {
		fmt.Println("error to connect to database")
		return
	}
	defer db.Close()

	// get query
	var queryFinal string
	query := "exec SP_DETAIL_PRODUCTION_LONGTERM " + " "
	where := "'" + order + "', '" + sort + "', '" + page + "', '" + pageSize + "', 'where a.ldc_id = ''" + ldc_id_param + "''" + " "

	// filter polis
	if noPolis != "" {
		andPolis := " and no_polis in (''" + noPolis + "'','''')"
		queryFinal = query + where + andPolis
	} else {
		queryFinal = query + where
	}

	// filter bisnis
	if business != "" {
		if beginDate == "" || endDate == "" {
			c.JSON(http.StatusOK, gin.H{ // nnti status ok nya di ganti
				"status":  400,
				"data":    "",
				"message": "failed get data, please provide valid date periode",
			})
			return
		}
		whereBusiness := " and LBU_NOTE like ''%" + business + "%'' "
		queryFinal = queryFinal + whereBusiness
	}

	// filter tgl
	if beginDate != "" && endDate != "" {
		whereDate := " and CAST(left(tgl_prod, 4) + right(TGL_PROD, 2) + left(right(TGL_PROD, 5), 2)  AS INT) between ''" + beginDate + "'' and ''" + endDate + "'' "
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
	tempFile := "produksi_longterm.xlsx" // Temporarily save the file
	if err := file.SaveAs(tempFile); err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Serve the Excel file as a response
	defer os.Remove(tempFile) // Remove the file after serving
	c.Header("Content-Disposition", "attachment; filename=produksi_longterm.xlsx")
	c.Header("Content-Type", "application/octet-stream")
	c.File(tempFile)
}
