package auth_temp

// ga kepake
// nnti di hapus
import (
	// "asm-backend/helper"

	// "asm-backend/wrapper"
	"asm-backend/auth"
	"asm-backend/model"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var authMiddleware *jwt.GinJWTMiddleware

func Login(c *gin.Context) {
	// Retrieve form values
	// case sensitive!!!!
	nik := c.PostForm("nik")
	password := c.PostForm("password")

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

	// cek if tgl akhir > current date
	// later

	// JWT PROCESS
	authMiddleware, err = model.JwtToken(nik, password, nikDb, passwordDb)

	if err != nil {
		fmt.Println("error on jwt : ", err)
	}

	authMiddleware.LoginHandler(c)
	// END JWT PROCESS
}

func RefreshLogin(c *gin.Context) {
	// JWT PROCESS
	// authMiddleware, err := helper.CurrentToken()
	// fmt.Println(err)

	if authMiddleware == nil {
		responseData := gin.H{
			"status": "400",
			"msg":    "token not found",
		}
		c.JSON(400, responseData)
	}

	authMiddleware.RefreshHandler(c)
}

func Logout(c *gin.Context) {
	// JWT PROCESS
	// authMiddleware, err := helper.CurrentToken()
	// fmt.Println(err)
	authMiddleware, err := auth.Token()

	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	if authMiddleware == nil {
		responseData := gin.H{
			"status": "400",
			"msg":    "token not found",
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
	// authMiddleware = nil
}

func CurrentToken() (*jwt.GinJWTMiddleware, error) {
	if authMiddleware == nil {
		return nil, nil
	}
	return authMiddleware, nil
}