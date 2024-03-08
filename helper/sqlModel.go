package helper

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

func SqlModel() (*sql.DB, error) {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Load env failed")
		// fmt.Println("Error connect to database")
		return nil, err
	}

	hostname := os.Getenv("hostname_sql")
	username := os.Getenv("username_sql")
	password := os.Getenv("password_sql")
	database := os.Getenv("database_sql")

	query := url.Values{}
	query.Add("database", database)

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(username, password),
		Host:     fmt.Sprintf("%s", hostname),
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("sqlserver", u.String())

	if err != nil {
		fmt.Println("error to connect to sql database!")
		fmt.Println("the error is : ", err)
		return nil, err
	}

	return db, nil
}
