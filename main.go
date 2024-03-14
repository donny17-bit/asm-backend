package main

import (
	"asm-backend/auth"
	"asm-backend/auth_token"
	"asm-backend/helper"
	"asm-backend/web"

	"fmt"
	"os"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()

	if err != nil {
		fmt.Print("Load env failed")
		return
	}

	authMiddleware, err := auth_token.Token()

	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	router.GET("/api/refresh", auth.RefreshLogin)
	router.GET("/api/logout", authMiddleware.LogoutHandler)
	router.POST("/api/login", authMiddleware.LoginHandler)

	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/production", web.Production)
		auth.GET("/hello", web.HelloHandler)
	}

	port := os.Getenv("PORT")
	fmt.Print("you are using port : ", port)
	router.Run(":" + port)
}

func Routing() *jwt.GinJWTMiddleware {
	authMiddleware, err := helper.JwtToken("test", "test", "test", "test")

	if err != nil {
		fmt.Println("terdapat error di authMiddleware")
		return nil
	}

	// if authMiddleware == nil {
	// 	fmt.Println("token not generated yet")
	// fmt.Println(err)
	// 	return nil
	// }
	return authMiddleware
}
