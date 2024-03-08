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

	port := os.Getenv("PORT")
	fmt.Print("you are using port : ", port)
	router.Run(":" + port)
}
