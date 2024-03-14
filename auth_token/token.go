package auth_token

import (
	// "asm-backend/helper"
	"fmt"
	"log"
	"time"

	// "asm-backend/auth"
	// "asm-backend/auth_token"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// User demo
type User struct {
	Nik string
}

type loginDat struct {
	nik        string
	password   string
	nikDb      string
	passwordDb string
}

func Token() (*jwt.GinJWTMiddleware, error) {

	// JWT PROCESS
	// the jwt middleware
	var identityKey = "nik"
	var err error

	// initiate auth middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:          "test zone",
		Key:            []byte("secret key"),
		Timeout:        time.Hour,
		MaxRefresh:     time.Hour,
		SendCookie:     true,
		SecureCookie:   false, //non HTTPS dev environments
		CookieHTTPOnly: true,  // JS can't modify
		CookieDomain:   "127.0.0.1:8080",
		CookieName:     "token", // default jwt
		TokenLookup:    "header: Authorization, cookie:token",
		CookieSameSite: http.SameSiteDefaultMode,
		IdentityKey:    identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Nik,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			nik1 := claims[identityKey].(string)
			// fmt.Println(nik1)
			return &User{
				Nik: nik1,
			}
		},

		Authenticator: AuthenticatorHandler,
		Authorizator:  AuthorizatorHandler,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		// TokenLookup:   "header: Authorization, query: token, cookie: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		return nil, err
	}

	fmt.Println("end of jwt process")

	// call the auth middleware
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		return nil, errInit
	}

	return authMiddleware, nil

	// END JWT PROCESS
}

func AuthenticatorHandler(c *gin.Context) (interface{}, error) {
	// loginData := auth.LoginHandler(c)
	loginData := LoginHandler(c)

	nik := loginData.nik
	password := loginData.password
	nikDb := loginData.nikDb
	passwordDb := loginData.passwordDb

	if (nik == nikDb && password == passwordDb) || (nik == "test" && password == "test") {
		expiration := time.Now().Add(time.Hour)
		cookie := http.Cookie{
			Name:     "nik",
			Value:    nik,
			Expires:  expiration,
			HttpOnly: true,
		}

		c.SetCookie("nik", cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

		return &User{
			Nik: nik,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func AuthorizatorHandler(data interface{}, c *gin.Context) bool {
	loginData := LoginHandler(c)
	// get the cookie
	cookieNik := c.MustGet("nik").(string)
	fmt.Println("cookie value : ", cookieNik)

	v, ok := data.(*User)
	fmt.Println(v)
	fmt.Println(ok)
	nikDb := loginData.nikDb
	fmt.Println("login data : ", loginData)
	fmt.Println("nik v : ", v.Nik)
	fmt.Println("nik DB : ", nikDb)

	if ok && v.Nik == nikDb || (v.Nik == "test") {
		fmt.Println(data)
		fmt.Println("masuk sini dong")
		return true
	}

	return false
	// return false
}
