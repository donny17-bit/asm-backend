package auth

import (
	"asm-backend/helper"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Print("Accessing login page")

	c.JSON(200, gin.H{
		"message": "This is the login function",
	})

	db, err := helper.OraModel()

	if err != nil {
		fmt.Println("error connecting to oracle database")
	}

	fmt.Println("success connecting to oracle database")

	// check for error
	err = db.Ping()
	// check for error

	if err != nil {
		fmt.Println("error pinging to oracle database")
	}
}
