package web

import (
	"asm-backend/helper"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Production(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "This is the production function",
	})

	db, err := helper.SqlModel()

	if err != nil {
		fmt.Println("error to connect to database")
		return
	}

	fmt.Println("connected to sql database")

	defer db.Close()

	// Perform a query
	rows, err := db.Query("SELECT top 10 no_polis, tgl_prod, prod_ke FROM warehouse_asm.dbo.production")
	fmt.Println(rows)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close() // Close the result set when done

	for rows.Next() {
		var (
			column1 string
			column2 string
			column3 string
		)
		err := rows.Scan(&column1, &column2, &column3)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		// Do something with the row data
		fmt.Println(column1, column2, column3)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}
}
