package web

import (
	// "asm-backend/auth_temp"
	"asm-backend/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HelloHandlerSession(c *gin.Context) {
	session := sessions.Default(c)

	nik := session.Get("nik")
	lastActivity := session.Get("lastActivity")

	if nik == nil || lastActivity == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"nik":          nik,
			"lastActivity": lastActivity,
			"text":         "unauthorized",
		})
		return
	}

	fmt.Println("current time - last activity", time.Now().Unix()-lastActivity.(int64))

	// Check if session is still active
	if time.Now().Unix()-lastActivity.(int64) > 1800 { // 30 minutes
		session.Options(sessions.Options{
			MaxAge: 0, // 0 minutes
		})
		session.Save()
		session.Clear()
		session.Save()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Session expired"})
		return
	}

	// Update last activity
	session.Set("lastActivity", time.Now().Unix())
	session.Options(sessions.Options{
		MaxAge: 1800, // 30 minutes
	})
	session.Save()

	c.JSON(200, gin.H{
		"nik":          nik,
		"lastActivity": lastActivity,
		"text":         "authorized",
	})
}

func Production(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "This is the production function",
	})

	db, err := model.SqlModel()

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
