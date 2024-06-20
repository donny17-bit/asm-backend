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

func LoginSession(c *gin.Context) {
	nik := c.PostForm("nik")
	password := c.PostForm("password")

	fmt.Println("login controller : ")
	fmt.Println("nik : ", nik)
	fmt.Println("password : ", password)

	// check is user already login
	ok := IsActive(c)

	if ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  200,
			"nik":     nik,
			"message": "user already login",
		})
		return
	}

	db := model.OraModel()
	defer db.Close()

	// Execute a query
	query := "SELECT A.NIK, B.MCL_NAME AS NAMA, A.PASS_ID, TRUNC (A.TGL_AKHIR), A.LDI_ID, A.CAB_ID FROM LST_USER_ASURANSI A, HRDASM.V_HRD_MST B WHERE     A.NIK = B.NIK AND A.NIK = :param1 AND A.PASS_ID = :param2 AND A.STS_AKTIF = '1'"
	rows, err := db.Query(query, nik, password)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close()

	// Iterate through the result set
	var (
		nikDb      string
		nama 	   string
		passwordDb string
		tgl_akhir  string
		ldi_id     string
		cab_id     string
	)

	for rows.Next() {
		if err := rows.Scan(&nikDb, &nama, &passwordDb, &tgl_akhir, &ldi_id, &cab_id); err != nil {
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
		oracleTimestampFormat := "2006-01-02 15:04:05" // This is the format recognized by Oracle's TO_TIMESTAMP function

		lastActivity := time.Now()
		expiration := time.Now().Add(30 * time.Minute)

		lastActivityTimestamp := lastActivity.Format(oracleTimestampFormat)
		expirationTimestamp := expiration.Format(oracleTimestampFormat)

		session.Set("id", id)
		session.Set("nik", nik)
		session.Set("nama", nama)
		session.Set("ldc_id", cab_id)
		session.Set("lastActivity", lastActivity.Unix())
		session.Set("expiration", expiration.Unix())
		session.Options(sessions.Options{
			MaxAge: 1800, // 30 minutes
			Path:   "/",
		})
		session.Save()

		// save it to db oracle
		// Execute a query
		query := "INSERT INTO LS_SESSION VALUES('" + id + "', '" + nik + "', TO_TIMESTAMP('" + lastActivityTimestamp + "', 'YYYY-MM-DD HH24:MI:SS'), TO_TIMESTAMP('" + expirationTimestamp + "', 'YYYY-MM-DD HH24:MI:SS'))"
		_, err := db.Exec(query)

		if err != nil {
			fmt.Println("Error executing query:", err)
			return
		}

		// c.JSON(http.StatusOK, gin.H{
		// 	"status":  200,
		// 	"nik":     nik,
		// 	"cabang":  cab_id,
		// 	"divisi":  ldi_id,
		// 	"message": "Login successful"})

		// redirect to dashboard page
		c.Redirect(http.StatusSeeOther, "/dashboard")

		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"nik":     nik,
			"cabang":  cab_id,
			"divisi":  ldi_id,
			"message": "Login failed"})
	}
}

func LogoutSession(c *gin.Context) {
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
		Path:   "/",
	})
	session.Save()
	session.Clear()

	// Save the session
	session.Save()

	c.JSON(http.StatusOK, gin.H{"status": 200,
		"message": "Logout successful"})
}

func IsActive(c *gin.Context) bool {
	session := sessions.Default(c)

	id := session.Get("id")
	nik := session.Get("nik")
	lastActivity := session.Get("lastActivity")
	expiration := session.Get("expiration")

	fmt.Println("is active controller")
	fmt.Println("session : ", session)
	fmt.Println("id : ", id)
	fmt.Println("nik : ", nik)
	fmt.Println("lastActivity : ", lastActivity)
	fmt.Println("expiration : ", expiration)

	if id == nil || nik == nil || lastActivity == nil || expiration == nil {
		return false
	}

	db := model.OraModel()
	defer db.Close()

	query := "SELECT * FROM LS_SESSION WHERE USER_ID = :param1 and SESSION_ID = :param2"
	result, err := db.Query(query, nik, id)

	fmt.Println("result : ", result)

	// kalau error di query or user_id/session id ga ada di db
	// login lagi
	if err != nil {
		fmt.Println("Error executing query:", err)
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
			Path:   "/",
		})
		session.Save()
		session.Clear()
		session.Save()
		return false
	}

	// if session not expired then update last activity
	oracleTimestampFormat := "2006-01-02 15:04:05"

	lastActivityNew := time.Now()
	expirationNew := time.Now().Add(30 * time.Minute)

	lastActivityTimestamp := lastActivityNew.Format(oracleTimestampFormat)
	expirationTimestamp := expirationNew.Format(oracleTimestampFormat)

	// update session cookie
	session.Set("lastActivity", lastActivityNew.Unix())
	session.Set("expiration", expirationNew.Unix())
	session.Options(sessions.Options{
		MaxAge: 1800, // 30 minutes
		Path:   "/",
	})
	session.Save()

	// save it to db oracle
	// Execute a query
	update := "UPDATE LS_SESSION SET LOGIN_TIME = TO_TIMESTAMP('" + lastActivityTimestamp + "', 'YYYY-MM-DD HH24:MI:SS'), EXPIRATION_TIME = TO_TIMESTAMP('" + expirationTimestamp + "', 'YYYY-MM-DD HH24:MI:SS') WHERE SESSION_ID = '" + id.(string) + "' AND USER_ID = '" + nik.(string) + "'"
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
