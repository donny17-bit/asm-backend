package main

import (
	"asm-backend/controller/auth"
	"asm-backend/controller/master"
	"asm-backend/controller/production"
	"asm-backend/views"
	"asm-backend/web"
	"fmt"
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

	// load all the static file
	router.LoadHTMLGlob("views/*.html")
	router.Static("/static/css", "./static/css")
	router.Static("/static/fonts", "./static/fonts")
	router.Static("/static/image", "./static/image")
	router.Static("/static/js", "./static/js")
	router.Static("/static/scss", "./static/scss")
	router.Static("/static/vendor", "./static/vendor")

	// views
	router.GET("/login", views.GetLogin)
	router.GET("/dashboard", views.GetDashboard)
	router.GET("/production-longterm", views.GetProductionLt)
	router.GET("/production-yearly", views.GetProductionYr)

	// production lt
	router.POST("/api/production-longterm", production.ProductionLt)
	router.POST("/api/export-production-longterm", production.ExportProdLt)

	// production yr
	router.POST("/api/production-yearly", production.ProductionYr)
	router.POST("/api/export-production-yearly", production.ExportProdYr)

	// master
	router.GET("/api/branch", master.GetBranch)
	router.GET("/api/business", master.GetBusiness)
	router.GET("/api/group-business", web.GetGrpBusiness)
	router.GET("/api/business-source", web.GetBusinessSource)

	// auth
	// router.POST("/api/login", auth.LoginSessionSql) // sql
	router.POST("/api/login", auth.LoginSession) // oracle
	// router.POST("/api/user-division", auth.LoginSession) // blm dipake
	router.GET("/api/logout", auth.LogoutSessionSql)

	// nnti diganti kalo udh login redirect langsung ke dalam menu dasboard
	router.GET("/", views.GetLogin)

	port := os.Getenv("PORT")
	fmt.Print("you are using port : ", port)
	router.Run(":" + port)
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Signature, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		// c.Writer.Header().Set("Content-Type", "application/json, text/html")
		// c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		// c.Writer.Header().Set("X-Frame-Options", "SAMEORIGIN")
		// c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
