package web

import (

	// "net/http"

	"github.com/gin-gonic/gin"
)

func Production(c *gin.Context) {
	
	c.JSON(200, gin.H{
		"message": "This is the production function",
	})
}
