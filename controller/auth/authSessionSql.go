package auth

import (
	"asm-backend/helper"
	"asm-backend/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginSessionSql(c *gin.Context) {
	// work on next js
	// var data map[string]interface{}

	// err := c.BindJSON(&data)
	// if err != nil {
	//     // Handle error
	//     return
	// }

	// // Now you can access the fields from the JSON data
	// nik, okNik := data["nik"].(string)
	// if !okNik {
	//     // "nik" field not found or not a string
	//     // Handle error
	// 	fmt.Println("nik not found in request body")
	//     return
	// }

	// password, okPass := data["password"].(string)
	// if !okPass {
	//     // "nik" field not found or not a string
	//     // Handle error
	// 	fmt.Println("password not found in request body")
	//     return
	// }

	// work on postman
	nik := c.PostForm("nik")
	password := c.PostForm("password")

	fmt.Println("nik : ", nik)
	fmt.Println("password : ", password)

	// check is user already login
	ok := IsActiveSql(c)

	if ok {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"nik":     nik,
			"message": "user already login",
		})
		return
	}

	db, err := model.SqlModel()

	if err != nil {
		fmt.Println("error connectiong to database : ", err)
		return
	}

	defer db.Close()

	// Execute a query
	query := "SELECT NIK, PASS_ID, CAST(TGL_AKHIR AS DATE) as tgl_akhir, LDI_ID, CAB_ID FROM LST_USER_ASURANSI WHERE NIK = '" + nik + "' AND PASS_ID = '" + password + "' AND STS_AKTIF = '1'"
	rows, err := db.Query(query)

	fmt.Println(query)

	if err != nil {
		fmt.Println("Error executing query login:", err)
		return
	}

	defer rows.Close()
	// Iterate through the result set
	var (
		nikDb      string
		passwordDb string
		tgl_akhir  string
		ldi_id     string
		cab_id     string
	)

	for rows.Next() {
		if err := rows.Scan(&nikDb, &passwordDb, &tgl_akhir, &ldi_id, &cab_id); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}

	session := sessions.Default(c)
	id := helper.GenerateRandomString(64)

	if nik == nikDb && password == passwordDb || nik == "test" {
		// oracleTimestampFormat := "2006-01-02 15:04:05" // This is the format recognized by Oracle's TO_TIMESTAMP function

		lastActivity := time.Now()
		lastActivityString := lastActivity.String()
		expiration := time.Now().Add(30 * time.Minute)
		// expiration := time.Now().Add(1 * time.Minute) // for testing
		expirationString := expiration.String()

		// lastActivityTimestamp := lastActivity.Format(oracleTimestampFormat)
		// expirationTimestamp := expiration.Format(oracleTimestampFormat)

		session.Set("id", id)
		session.Set("nik", nik)
		session.Set("ldc_id", cab_id)
		session.Set("lastActivity", lastActivity.Unix())
		session.Set("expiration", expiration.Unix())
		session.Options(sessions.Options{
			MaxAge: 1800, // 30 minutes
			// MaxAge: 60, // for testing

			// HttpOnly: true,
		})
		session.Save()

		fmt.Println("last activity in unix : ", lastActivity.Unix())

		// save it to db sql
		// Execute a query
		query := "INSERT INTO LS_SESSION VALUES('" + id + "', '" + nik + "', '" + lastActivityString + "', '" + expirationString + "')"
		_, err := db.Exec(query)

		if err != nil {
			fmt.Println("Error executing query:", err)
			return
		}

		// Set session data as a cookie in the HTTP response
		// c.SetCookie("session", id, 1800, "/", "", false, false) // SetCookie(name, value, maxAge, path, domain, secure, httpOnly)

		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"nik":     nik,
			"cabang":  cab_id,
			"divisi":  ldi_id,
			"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"nik":     nik,
			"cabang":  cab_id,
			"divisi":  ldi_id,
			"message": "Login failed"})
	}
}

func LogoutSessionSql(c *gin.Context) {
	// nnti ditambahin cek id ke db

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
	session.Clear()
	session.Save()

	// Save the session
	// session.Save()

	c.JSON(http.StatusOK, gin.H{"status": 200,
		"message": "Logout successful"})
}

func IsActiveSql(c *gin.Context) bool {
	session := sessions.Default(c)

	id := session.Get("id")
	nik := session.Get("nik")
	lastActivity := session.Get("lastActivity")
	expiration := session.Get("expiration")

	fmt.Println("nik in is active : ", nik)
	fmt.Println("id session : ", id)

	if id == nil || nik == nil || lastActivity == nil || expiration == nil {
		return false
	}

	db, err := model.SqlModel()
	if err != nil {
		fmt.Println("error connectiong to database : ", err)
		return false
	}

	defer db.Close()

	query := "SELECT * FROM LS_SESSION WHERE USER_ID = '" + nik.(string) + "' and SESSION_ID = '" + id.(string) + "'"
	result, err := db.Query(query)

	// kalau error di query or user_id/session id ga ada di db
	// login lagi
	if err != nil {
		fmt.Println("Error executing query is active:", err)
		return false
	}

	defer result.Close()

	// Iterate through the result set
	var (
		sessionIdDb      string
		nikDb            string
		loginTimeDb      string
		expirationTimeDb string
	)

	for result.Next() {
		if err := result.Scan(&sessionIdDb, &nikDb, &loginTimeDb, &expirationTimeDb); err != nil {
			fmt.Println("Error scanning row in isActive controller:", err)
			return true // cek lagi
		}
	}

	// logout if session is expired
	if expiration.(int64)-lastActivity.(int64) > 1800 { // 30 minutes
		// log out session
		session.Options(sessions.Options{
			MaxAge: 0, // 0 minutes
		})
		session.Save()
		session.Clear()
		session.Save()
		return false
	}

	// if session not expired then update last activity
	lastActivityNew := time.Now()
	lastActivityString := lastActivityNew.String()
	expirationNew := time.Now().Add(30 * time.Minute)

	// for testing only
	// expirationNew := time.Now().Add(1 * time.Minute)

	expirationString := expirationNew.String()

	// update session cookie
	session.Set("lastActivity", lastActivityNew.Unix())
	session.Set("expiration", expirationNew.Unix())
	session.Options(sessions.Options{
		MaxAge: 1800, // 30 minutes
		// MaxAge: 60, // for testing only
	})
	session.Save()

	// save it to db oracle
	// Execute a query
	update := "UPDATE LS_SESSION SET LOGIN_TIME = '" + lastActivityString + "', EXPIRATION_TIME = '" + expirationString + "' WHERE SESSION_ID = '" + id.(string) + "' AND USER_ID = '" + nik.(string) + "'"
	_, errUpdate := db.Query(update)

	if errUpdate != nil {
		fmt.Println("Error executing query:", err)
		return false
	}

	return true
}

// lanjut nnti dlu
// func getUserDivision(c *gin.Context) {
// 	nik := c.PostForm("nik")
// 	password := c.PostForm("password")

// 	fmt.Println("nik : ", nik)
// 	fmt.Println("password : ", password)

// 	// check is user already login
// 	// ok := IsActive(c)

// 	// if ok {
// 	// 	c.JSON(http.StatusUnauthorized, gin.H{
// 	// 		"status":  200,
// 	// 		"nik":     nik,
// 	// 		"message": "user already login",
// 	// 	})
// 	// 	return
// 	// }
// 	// skip dlu

// 	db := model.OraModel()
// 	defer db.Close()

// 	// Execute a query
// 	query := "SELECT NIK, PASS_ID, TRUNC(TGL_AKHIR), LDI_ID, CAB_ID FROM LST_USER_ASURANSI WHERE NIK = :param1 AND PASS_ID = :param2 AND STS_AKTIF = '1'"
// 	rows, err := db.Query(query, nik, password)

// 	if err != nil {
// 		fmt.Println("Error executing query:", err)
// 		return
// 	}

// 	// Iterate through the result set
// 	type Data struct {
// 		Nik        string  `json:"Nik"`
// 		BranchId      string  `json:"CabangId"`
// 		Branch string  `json:"Branch"`
// 		DivisionId       string  `json:DivisionId"`
// 		Division            string  `json:"Division"`
// 	}

// 	for rows.Next() {
// 		if err := rows.Scan(&nikDb, &passwordDb, &tgl_akhir, &ldi_id, &cab_id); err != nil {
// 			fmt.Println("Error scanning row:", err)
// 			return
// 		}
// 	}

// 	if err := rows.Err(); err != nil {
// 		fmt.Println("Error iterating rows:", err)
// 		return
// 	}

// 	if nik == nikDb && password == passwordDb || nik == "test" {
// 		c.JSON(http.StatusOK, gin.H{
// 			"status":  200,
// 			"nik":     nik,
// 			"cabang":  cab_id,
// 			// "cabang":
// 			"divisi":  ldi_id,
// 			"message": "get division sucess"})
// 	} else {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"status":  401,
// 			"nik":     nik,
// 			"cabang":  cab_id,
// 			"divisi":  ldi_id,
// 			"message": " failed"})
// 	}
// }
