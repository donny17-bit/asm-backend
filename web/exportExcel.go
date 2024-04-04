package web

import (
	"asm-backend/auth"
	"asm-backend/model"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func exportProdLt(c *gin.Context) {
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
	

}