package main

import (
	"asm-backend/auth"
	"asm-backend/web"

	"fmt"
	"os"

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

	authMiddleware, err := auth.Token()

	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	router.GET("/api/refresh", auth.RefreshToken)
	router.GET("/api/logout", auth.Logout)
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
