package auth

import (
	"asm-backend/model"
	"fmt"
	"net/http"
	"time"

	// jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	// "github.com/gorilla/sessions"
)

type loginData struct {
	nik        string
	password   string
	nikDb      string
	passwordDb string
}

func Login(c *gin.Context) loginData {
	// Retrieve form values
	// case sensitive!!!!
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
		return loginData{}
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
			return loginData{}
		}

		// Do something with the row data
		fmt.Println(nik, password, tgl_akhir, ldi_id)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return loginData{}
	}

	login := loginData{}
	login.nik = nik
	login.password = password
	login.nikDb = nikDb
	login.passwordDb = passwordDb

	return login

	// cek if tgl akhir > current date
	// later
}

func RefreshToken(c *gin.Context) {
	authMiddleware, err := Token()

	if err != nil || authMiddleware == nil {
		responseData := gin.H{
			"status": "400",
			"msg":    "token not found",
			"error":  err,
		}
		c.JSON(400, responseData)
	}

	authMiddleware.RefreshHandler(c)
}

func Logout(c *gin.Context) {
	authMiddleware, err := Token()

	if err != nil || authMiddleware == nil {
		responseData := gin.H{
			"status": "400",
			"msg":    "token not found",
			"error":  err,
		}
		c.JSON(400, responseData)
	}

	cookie := http.Cookie{
		Name:     "nik",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}

	c.SetCookie("nik", cookie.Value, 0, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	authMiddleware.LogoutHandler(c)
}