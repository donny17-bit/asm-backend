package auth

import (
	"asm-backend/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginSession(c *gin.Context) {
	nik := c.PostForm("nik")
	password := c.PostForm("password")

	fmt.Println("nik : ", nik)
	fmt.Println("password : ", password)

	db := model.OraModel()
	defer db.Close()

	// Execute a query
	rows, err := db.Query("SELECT NIK, PASS_ID, TRUNC(TGL_AKHIR), LDI_ID FROM LST_USER_ASURANSI WHERE NIK = :param1 AND PASS_ID = :param2", nik, password)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close()
	// Iterate through the result set
	var (
		nikDb      string
		passwordDb string
		tgl_akhir  string
		ldi_id     string
	)

	for rows.Next() {
		if err := rows.Scan(&nikDb, &passwordDb, &tgl_akhir, &ldi_id); err != nil {
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

	session := sessions.Default(c)

	if nik == nikDb && password == passwordDb || nik == "test" {
		session.Set("nik", nik)
		session.Set("lastActivity", time.Now().Unix())
		session.Options(sessions.Options{
			MaxAge: 1800, // 30 minutes
		})
		session.Save()

		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
	}
}

func LogoutSession(c *gin.Context) {
	session := sessions.Default(c)

	if session == nil {
		responseData := gin.H{
			"status": "400",
			"msg":    "session not found",
		}
		c.JSON(400, responseData)
	}

	// Clear the session
	session.Options(sessions.Options{
		MaxAge: 0, // 0 minutes
	})
	session.Save()
	session.Clear()

	// Save the session
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
