package master

import (
	"asm-backend/controller/auth"
	"fmt"
	"net/http"

	"asm-backend/model"

	"github.com/gin-gonic/gin"
)

func GetBusiness(c *gin.Context) {

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
	query := "select * from lst_business order by lbu_note"

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close() // Close the result set when done

	type Data struct {
		LbuId    string
		LgbId    string
		LbuNote  string
		TypeNote string
	}

	var datas []Data

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.LbuId, &data.LgbId, &data.LbuNote, &data.TypeNote)

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
