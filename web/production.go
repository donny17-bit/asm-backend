package web

import (
	"asm-backend/auth"
	"asm-backend/helper"
	"fmt"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func ProductionAuth(router *gin.Engine) {
	authMiddleware, err := auth.CurrentToken()

	if err != nil {
		fmt.Println("terdapat error di authMiddleware")
		// c.JSON(200, gin.H{
		// 	"message": "no auth middleware",
		// })
	}

	if authMiddleware == nil {
		fmt.Println("authMiddleware is nil")
		return
		// c.JSON(200, gin.H{
		// 	"message": "no auth middleware",
		// })
	}

	// authMiddleware.MiddlewareFunc()
	// app := router.Group("/")
	router.Use(authMiddleware.MiddlewareFunc())
	{
		router.GET("/productionAuth", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "authenticated middleware",
			})
		})
	}

}

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	// fmt.Println(claims)
	user, _ := c.Get("nik")

	fmt.Println("isi user : ", user)
	c.JSON(200, gin.H{
		"nik": claims["nik"],
		//   "userName": user.(*User).nik,
		"text": "Hello World.",
	})
}

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
