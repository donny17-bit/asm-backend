package web

import (
	"asm-backend/auth"
	"fmt"
	"net/http"

	"asm-backend/model"

	"github.com/gin-gonic/gin"
)

func GetGrpBusiness(c *gin.Context) {

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

	// get query
	query := "select lgb_id, lgb_note from lst_grp_business order by lgb_note"

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close() // Close the result set when done

	type Data struct {
		LgbId   string
		LgbNote string
	}

	var datas []Data

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.LgbId, &data.LgbNote)

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
