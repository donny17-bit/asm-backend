package web

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Production(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "This is the production function",
	})

	//sqlserver://sa:mypass@localhost:1234?database=master&connection+timeout=30

	hostname := "192.168.108.63"
	username := "sa"
	password := "AdminSQL2008#!"

	query := url.Values{}
	query.Add("database", "Warehouse_ASM")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(username, password),
		Host:   fmt.Sprintf("%s", hostname),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}

	fmt.Print(u.String())
	db, err := sql.Open("sqlserver", u.String())

	if err != nil {
		fmt.Print("database sql error connected")
	} else {
		fmt.Print("database sql connected")
	}

	fmt.Print(db)

}
