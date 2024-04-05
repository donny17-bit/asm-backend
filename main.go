package main

import (
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

// use session
func main() {
	router := gin.Default()
	router.Use(CORS())
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
	router.GET("/api/production-yearly", web.GetProductionYr)
	router.GET("/api/export-production-longterm", web.ExportProdLt)

	// master
	router.GET("/api/branch", web.GetBranch)
	router.GET("/api/business", web.GetBusiness)
	router.GET("/api/group-business", web.GetGrpBusiness)
	router.GET("/api/business-source", web.GetBusinessSource)

	// auth
	router.POST("/api/login", auth.LoginSessionSql) // sql
	// router.POST("/api/login", auth.LoginSession) // oracle
	// router.POST("/api/user-division", auth.LoginSession) // blm dipake
	router.GET("/api/logout", auth.LogoutSessionSql)

	port := os.Getenv("PORT")
	fmt.Print("you are using port : ", port)
	router.Run(":" + port)
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Signature, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-Frame-Options", "SAMEORIGIN")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
