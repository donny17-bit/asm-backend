package main

import (
	// "asm-backend/auth"
	// "asm-backend/web"
	"asm-backend/auth"
	"asm-backend/web"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// use jwt token
// func main() {
// 	router := gin.Default()

// 	err := godotenv.Load()

// 	if err != nil {
// 		fmt.Print("Load env failed")
// 		return
// 	}

// 	authMiddleware, err := auth.Token()

// 	if err != nil {
// 		fmt.Println("err : ", err)
// 		return
// 	}

// 	router.GET("/api/refresh", auth.RefreshToken)
// 	router.GET("/api/logout", auth.Logout)
// 	router.POST("/api/login", authMiddleware.LoginHandler)

// 	auth := router.Group("/auth")
// 	auth.Use(authMiddleware.MiddlewareFunc())
// 	{
// 		auth.GET("/production", web.Production)
// 		auth.GET("/hello", web.HelloHandler)
// 	}

// 	port := os.Getenv("PORT")
// 	fmt.Print("you are using port : ", port)
// 	router.Run(":" + port)
// }

// use session
func main() {
	router := gin.Default()

	err := godotenv.Load()

	if err != nil {
		fmt.Print("Load env failed")
		return
	}

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	// main route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "base route",
		})
	})

	// production
	router.GET("/api/production-longterm", web.GetProductionLt)
	router.GET("/api/production-yearly", web.ProductionYr)

	// auth
	router.POST("/api/login", auth.LoginSession)
	router.POST("/api/user-division", auth.LoginSession) // blm dipake
	router.GET("/api/logout", auth.LogoutSession)

	port := os.Getenv("PORT")
	fmt.Print("you are using port : ", port)
	router.Run(":" + port)
}
