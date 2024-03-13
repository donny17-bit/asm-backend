package main

import (
	"asm-backend/auth"
	"asm-backend/helper"
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

	authMiddleware, err := helper.CurrentToken()
	if authMiddleware == nil {
		fmt.Println("token not generated yet")
		fmt.Println(err)
	}

	// routing
	router.Use(authMiddleware.MiddlewareFunc())
	{
		router.GET("/api/production", web.Production)
	}
	// router.GET("/api/production", web.Production)
	router.GET("/api/refresh", auth.RefreshLogin)
	router.GET("/api/logout", auth.Logout)

	router.POST("/api/login", auth.Login)
	// router.POST("/api/login", auth.Jwt().LoginHandler)

	port := os.Getenv("PORT")
	fmt.Print("you are using port : ", port)
	router.Run(":" + port)
}
