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

	// routing
	router.GET("/api/production", web.Production)
	router.GET("/api/login", auth.Login)
	router.GET("/api/refresh", auth.RefreshLogin)

	router.POST("/api/login", auth.PostLogin)
	// router.POST("/api/login", auth.Jwt().LoginHandler)

	port := os.Getenv("PORT")
	fmt.Print("you are using port : ", port)
	router.Run(":" + port)
}
