package web

import (
	"asm-backend/auth"
	"asm-backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Production(c *gin.Context) {

	ok := auth.IsActive(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	db, err := model.SqlModel()

	if err != nil {
		fmt.Println("error to connect to database")
		return
	}
	defer db.Close()

	// Perform a query
	rows, err := db.Query("SELECT top 10 no_polis, tgl_prod, prod_ke FROM warehouse_asm.dbo.production")
	fmt.Println(rows)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close() // Close the result set when done

	type Data struct {
		No_polis string `json:"no_polis"`
		Tgl_prod string `json:"tgl_prod"`
		Prod_ke  string `json:"prod_ke"`
	}

	// Create an array to store the query results
	var datas []Data

	for rows.Next() {
		var data Data

		// Scan each row into a struct
		if err := rows.Scan(&data.No_polis, &data.Tgl_prod, &data.Prod_ke); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// err := rows.Scan(&no_polis, &tgl_prod, &prod_ke)
		// if err != nil {
		// 	fmt.Println("Error scanning row:", err)
		// 	return
		// }

		// Append the struct to the array
		datas = append(datas, data)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}

	// if no error
	// Respond with JSON data
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   datas,
	})
}
