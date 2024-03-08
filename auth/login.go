package auth

import (
	"asm-backend/helper"
	// "asm-backend/wrapper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Print("Accessing login page")

	c.JSON(200, gin.H{
		"message": "This is the login function",
	})

	db, err := helper.OraModel()

	if err != nil {
		fmt.Println("error connecting to oracle database")
	}

	fmt.Println("success connecting to oracle database")

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

	db, err := helper.OraModel()

	if err != nil {
		fmt.Println("error occured when connecting to oracle database : ", err)
		return
	}

	defer db.Close()

	// Execute a query
	rows, err := db.Query("SELECT LUS_ID, NIK, PASS_ID FROM LST_USER_ASURANSI WHERE NIK = :param1 AND PASS_ID = :param2", nik, password)

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

	responseData := gin.H{
		"status": http.StatusOK,
		"msg":    "success get data",
		"data": []gin.H{
			{
				"email":    nik,
				"password": password,
			},
		},
	}

	c.JSON(http.StatusOK, responseData)

	// wrapper nnti aja
	// wrapper.JsonWrapper(http.StatusOK, c, responseData)
}
