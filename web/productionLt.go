package web

import (
	"asm-backend/auth"
	"asm-backend/model"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetProductionLt(c *gin.Context) {
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
	clientName := c.Query("client_name")

	if sort == "" {
		sort = "asc"
	}

	db, err := model.SqlModel()

	if err != nil {
		fmt.Println("error to connect to database")
		return
	}
	defer db.Close()

	// cek total row ------------------
	var queryRow string
	queryTotalRow := "select count(1) as total_rows FROM PRODUCTION_GABUNGAN_VIEW A JOIN MV_AGEN MA ON A.LAG_ID = MA.LAG_AGEN_ID JOIN LST_CABANG LC ON A.LDC_ID = LC.LDC_ID JOIN LST_BUSINESS LB ON A.LBU_ID = LB.LBU_ID JOIN LST_GRP_BUSINESS LGB ON LB.LGB_ID = LGB.LGB_ID JOIN LST_JN_PRODUKSI LJP ON LJP.LJP_ID = A.LJP_ID JOIN JNNER JNN ON JNN.JN_NER = A.JN_NER LEFT OUTER JOIN LST_MO MO ON A.LMO_ID = MO.LMO_ID LEFT OUTER JOIN MST_CLIENT MC ON A.CLIENT_ID = MC.MCL_ID LEFT OUTER JOIN LST_JENIS_COAS JN_COAS ON A.MDS_JN_COAS = JN_COAS.MDS_JN_COAS "
	whereRow := "where a.ldc_id = '" + ldc_id_param + "' "

	// filter polis
	if noPolis != "" {
		andPolis := " and no_polis in ('" + noPolis + "','') "
		queryRow = queryTotalRow + whereRow + andPolis
	} else {
		queryRow = queryTotalRow + whereRow
	}

	// filter bisnis
	if business != "" {
		if beginDate == "" || endDate == "" {
			c.JSON(http.StatusOK, gin.H{ // nnti status ok nya di ganti status failed
				"status":  400,
				"data":    "",
				"message": "failed get data, please provide valid date periode",
			})
			return
		}
		whereBusiness := " and LBU_NOTE like '%" + business + "%' "
		queryRow = queryRow + whereBusiness
	}

	// filter client name
	if clientName != "" {
		if beginDate == "" || endDate == "" {
			c.JSON(http.StatusOK, gin.H{ // nnti status ok nya di ganti status failed
				"status":  400,
				"data":    "",
				"message": "failed get data, please provide valid date periode",
			})
			return
		}
		whereClient := " and MC.MCL_NAME like '%" + clientName + "%' "
		queryRow = queryRow + whereClient
	}

	// filter tgl
	if beginDate != "" && endDate != "" {
		whereDate := " and CAST(left(tgl_prod, 4) + right(TGL_PROD, 2) + left(right(TGL_PROD, 5), 2)  AS INT) between  '" + beginDate + "' and '" + endDate + "' "
		queryRow = queryRow + whereDate
	}

	countRows, err := db.Query(queryRow)

	if err != nil {
		fmt.Println("Error executing query row:", err)
		return
	}

	defer countRows.Close()

	var totalRows int
	// var totalPage float64

	for countRows.Next() {
		err := countRows.Scan(&totalRows)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error in count rows": err.Error()})
		}
	}

	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil {
		fmt.Println("Error convert page size to int :", err)
		return
	}

	totalPage := math.Ceil(float64(totalRows) / float64(pageSizeNum))
	// totalPageFloat := float64(totalPage)

	// end count rows
	// -----------------

	if err := countRows.Err(); err != nil {
		fmt.Println("Error iterating count rows:", err)
		return
	}

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

	// filter client name
	if clientName != "" {
		if beginDate == "" || endDate == "" {
			c.JSON(http.StatusOK, gin.H{ // nnti status ok nya di ganti
				"status":  400,
				"data":    "",
				"message": "failed get data, please provide valid date periode",
			})
			return
		}
		whereClient := " and MCL_NAME like ''%" + clientName + "%'' "
		queryFinal = queryFinal + whereClient
	}

	// filter tgl
	if beginDate != "" && endDate != "" {
		whereDate := " and CAST(left(tgl_prod, 4) + right(TGL_PROD, 2) + left(right(TGL_PROD, 5), 2)  AS INT) between ''" + beginDate + "'' and ''" + endDate + "'' "
		queryFinal = queryFinal + whereDate
	}

	queryFinal = queryFinal + "'"

	fmt.Println("queryFinal : ", queryFinal)
	// nnti tambhain klau yg login nik itasm, keluarin semua
	rows, err := db.Query(queryFinal)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close() // Close the result set when done

	
	type Data struct {
		Rn            string  `json:"Rn"`
		TglProd       string  `json:"TglProd"`
		ThnBln        string  `json:"ThnBln"`
		ProdDate      string  `json:"ProdDate"`
		BeginDate     string  `json:"BeginDate"`
		EndDate       string  `json:"EndDate"`
		Mo            string  `json:"Mo"`
		ClientName    string  `json:"ClientName"`
		Kanwil        string  `json:"Kanwil"`
		Cabang        string  `json:"Cabang"`
		Perwakilan    string  `json:"Perwakilan"`
		SubPerwakilan string  `json:"SubPerwakilan"`
		Jnner         string  `json:"Jnner"`
		JenisProd     string  `json:"JenisProd"`
		JenisPaket    *string `json:"JenisPaket"`
		// JenisPaket    NullableString `json:"JenisPaket"`
		Ket *string `json:"Ket"`
		// NamaCeding    NullableString `json:"NamaCeding"`
		NamaCeding    *string `json:"NamaCeding"`
		Namaleader0   string  `json:"Namaleader0"`
		Namaleader1   string  `json:"Namaleader1"`
		Namaleader2   string  `json:"Namaleader2"`
		Namaleader3   string  `json:"Namaleader3"`
		GroupBusiness string  `json:"GroupBusiness"`
		Business      string  `json:"Business"`
		NoKontrak     *string `json:"NoKontrak"`
		NoPolis       string  `json:"NoPolis"`
		NoCif         string  `json:"NoCif"`
		ProdKe        string  `json:"ProdKe"`
		NamaDealer    *string `json:"NamaDealer"`
		Tsi           float32 `json:"Tsi"`
		Gpw           float32 `json:"Gpw"`
		Disc          float32 `json:"Disc"`
		Disc2         float32 `json:"Disc2"`
		Comm          float32 `json:"Comm"`
		Oc            float32 `json:"Oc"`
		Bkp           float32 `json:"Bkp"`
		Ngpw          float32 `json:"Ngpw"`
		Ri            float32 `json:"Ri"`
		Ricom         float32 `json:"Ricom"`
		Npw           float32 `json:"Npw"`
	}

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
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			// fmt.Println("Error scanning row:", err)
		}

		// Append the struct to the array
		datas = append(datas, data)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		// Handle error
		fmt.Println("Error converting page to number:", err)
		return
	}
	nextPage := pageNum + 1
	previousPage := pageNum - 1

	// if no error
	// Respond with JSON data
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"data":          datas,
		"current_page":  page,
		"next_page":     nextPage,
		"previous_page": previousPage,
		"max_page":      totalPage,
		"page_size":     pageSize,
		"total_data":    totalRows,
		"message":       "success get data",
	})
}
