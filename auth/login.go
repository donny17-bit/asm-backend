package auth

import (
	"asm-backend/helper"

	// "asm-backend/wrapper"
	"fmt"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var authMiddleware *jwt.GinJWTMiddleware

func Login(c *gin.Context) {
	fmt.Print("Accessing login page")

	c.JSON(200, gin.H{
		"message": "This is the login function",
	})

	db := helper.OraModel()
	defer db.Close()

	// Ping the database to check the connection
	if err := db.Ping(); err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}

	// Execute a query
	rows, err := db.Query("SELECT LUS_ID, NIK, PASS_ID FROM LST_USER_ASURANSI WHERE ROWNUM < 10")
	// rows, err := db.QueryContext(context.Background(), "SELECT LUS_ID, NIK, PASS_ID FROM LST_USER_ASURANSI WHERE ROWNUM < 10")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var (
			column1 string
			column2 string
			column3 string
		)
		if err := rows.Scan(&column1, &column2, &column3); err != nil {
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

func PostLogin(c *gin.Context) {
	// Retrieve form values
	// case sensitive!!!!
	nik := c.PostForm("nik")
	password := c.PostForm("password")

	db := helper.OraModel()
	defer db.Close()

	// Execute a query
	rows, err := db.Query("SELECT NIK, PASS_ID, TRUNC(TGL_AKHIR), LDI_ID FROM LST_USER_ASURANSI WHERE NIK = :param1 AND PASS_ID = :param2", nik, password)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close()
	// Iterate through the result set
	for rows.Next() {
		var (
			nik       string
			password  string
			tgl_akhir string
			ldi_id    string
		)

		if err := rows.Scan(&nik, &password, &tgl_akhir, &ldi_id); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		// Do something with the row data
		fmt.Println(nik, password, tgl_akhir, ldi_id)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}

	// cek if tgl akhir > current date
	// later

	// JWT PROCESS
	authMiddleware, err = helper.JwtToken(nik, password)

	if err != nil {
		fmt.Println("error on jwt : ", err)
	}

	authMiddleware.LoginHandler(c)
	// END JWT PROCESS
}

func RefreshLogin(c *gin.Context) {
	// JWT PROCESS
	if authMiddleware == nil {
		responseData := gin.H{
			"status": "400",
			"msg":    "token not found",
		}
		c.JSON(400, responseData)
	}

	authMiddleware.RefreshHandler(c)
}
