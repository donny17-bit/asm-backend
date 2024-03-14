package helper

import (
	// "asm-backend/helper"
	"fmt"
	"log"
	"time"

	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// var authMiddleware *jwt.GinJWTMiddleware

func JwtToken(nik string, password string, nikDb string, passwordDb string) (*jwt.GinJWTMiddleware, error) {

	// JWT PROCESS
	// the jwt middleware
	var identityKey = "nik"
	var err error

	// User demo
	type User struct {
		nik string
	}

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
		// TokenLookup:      "cookie:token",
		CookieSameSite: http.SameSiteDefaultMode,
		IdentityKey:    identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.nik,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				nik: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {

			userNik := nik
			password := password

			if (userNik == nikDb && password == passwordDb) || (userNik == "test" && password == "test") {
				return &User{
					nik: userNik,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.nik == nikDb {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: token",
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

// balikin ke login package
// func CurrentToken() (*jwt.GinJWTMiddleware, error) {
// 	if authMiddleware == nil {
// 		return nil, nil
// 	}
// 	return authMiddleware, nil
// }
