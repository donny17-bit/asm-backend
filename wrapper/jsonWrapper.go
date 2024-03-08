package wrapper

import (
	"github.com/gin-gonic/gin"
)

// wrapper nnti aja
func JsonWrapper(status int, c *gin.Context, data any) {
	responseData := gin.H{
		"status": status,
		"msg":    data.msg,
		"data":   data,
		// "NIK":      nik,
		// "PASSWORD": password,
	}

	c.JSON(status, responseData)
}
