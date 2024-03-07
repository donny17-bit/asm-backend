package main

// import "fmt"
import (
	"asm-backend/web"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// router := gin.New()
	router := gin.Default()

	router.GET("/api/ping", web.Production)
	// Web.Production()
	// router.GET(apiRouter+"/Ping", WebAdmin.Article)

	// port := ":8080"
	port := os.Getenv("PORT")
	fmt.Print("you are using port : ", port)
	router.Run(":" + port)
}
