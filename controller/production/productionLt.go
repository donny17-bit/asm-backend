package production

import (
	"asm-backend/model"
	"fmt"

	// "math"
	"net/http"
	// "strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type InputData struct {
	Begin_date string `json:"begin_date"`
	End_date string `json:"end_date"`
	No_polis string `json:"no_polis"`
	No_cif string `json:"no_cif"`
	Client_name string `json:"Client_name"`
	Branch string `json:"Branch"`
	Business string `json:"business"`
	Sumbis string `json:"sumbis"`
}


func ProductionLt(c *gin.Context) {
	session := sessions.Default(c)
	ldc_id := session.Get("ldc_id") // default sesuai info login

	// if ldc_id == nil {
	// 	fmt.Println("error cabang kosong")
	// 	return
	// }

	var inputData InputData
	if err := c.BindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// // cek jika ldc_id ada di request
	var ldc_id_param string

	if ldc_id != nil {
		ldc_id_param = ldc_id.(string)
	} else {
		ldc_id_param = c.PostForm("ldc_id")
	}

	page := "1"          // req
	pageSize := "10"    // req
	sort := "asc"        // opt
	order := "thnbln, client_name"        // req
	noPolis := c.PostForm("no_polis")
	beginDate := inputData.Begin_date
	endDate := inputData.End_date
	business := c.PostForm("business")
	clientName := c.PostForm("client_name")

	var formatedBeginDate string
	var formatedEndDate string

	fmt.Println("ldc_id_param", ldc_id_param)
	fmt.Println("noPolis", noPolis)
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

	// cek total row ----------------------------------------------------------------
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

	countRows, err := db.Query(queryRow)

	if err != nil {
		fmt.Println("Error executing query row:", err)
		return
	}

	defer countRows.Close()

	var totalRows string

	// save query result to total rows variable
	for countRows.Next() {
		err := countRows.Scan(&totalRows)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error in count rows": err.Error()})
		}
	}

	fmt.Println("total rows", totalRows)

	// pageSizeNum, err := strconv.Atoi(pageSize)
	// if err != nil {
	// 	fmt.Println("Error convert page size to int :", err)
	// 	return
	// }

	// totalPage := math.Ceil(float64(totalRows) / float64(pageSizeNum))
	// // totalPageFloat := float64(totalPage)

	if err := countRows.Err(); err != nil {
		fmt.Println("Error iterating count rows:", err)
		return
	}

	// end count rows
	// ------------------------------------------------------------------------------------------------------------


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
		whereDate := " and CAST(left(tgl_prod, 4) + right(TGL_PROD, 2) + left(right(TGL_PROD, 5), 2)  AS INT) between ''" + formatedBeginDate + "'' and ''" + formatedEndDate + "'' "
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

	// pageNum, err := strconv.Atoi(page)
	// if err != nil {
	// 	// Handle error
	// 	fmt.Println("Error converting page to number:", err)
	// 	return
	// }
	// nextPage := pageNum + 1
	// previousPage := pageNum - 1

	// "data"	: datas, this is array object type

	// if no error
	// Respond with redirect html page
	// c.HTML(http.StatusOK, "production.html", gin.H{
	// 	"begin_date": beginDate,
	// 	"end_date": endDate,
	// 	"no_polis": noPolis,
	// 	"client_name": clientName,
	// 	"business": business,
	// 	"data"	: datas,
	// })

	// Respond with JSON data
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"data":          datas,
		"current_page":  page,
		// "next_page":     nextPage,
		// "previous_page": previousPage,
		// "max_page":      totalPage,
		"page_size":     pageSize,
		"total_data":    totalRows,
		"message":       "success get data",
	})
	
}
