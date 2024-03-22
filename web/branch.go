package web

import (
	"asm-backend/auth"
	"asm-backend/model"
	"fmt"
	"net/http"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetBranch(c *gin.Context) {

	ok := auth.IsActive(c)

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

	db, err := model.SqlModel()

	if err != nil {
		fmt.Println("error to connect to database")
		return
	}
	defer db.Close()

	// nnti kasih proteksi, hanya bisa filter pada cabang user tersebut
	// get query
	query := "select ldc_id, id, kanwil, cabang, perwakilan, sub_perwakilan from lst_cabang order by kanwil, cabang, perwakilan, sub_perwakilan"

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close() // Close the result set when done

	type Data struct {
		LdcId         string
		Id            string
		Kanwil        string
		Cabang        string
		Perwakilan    string
		SubPerwakilan string
	}

	var datas []Data

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.LdcId, &data.Id, &data.Kanwil, &data.Cabang, &data.Perwakilan, &data.SubPerwakilan)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		// Append the struct to the array
		datas = append(datas, data)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    datas,
		"message": "success get data",
	})
}
