package helper

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	go_ora "github.com/sijms/go-ora/v2"
)

func OraModel() (*sql.DB, error) {
	fmt.Println("this is ora model function")

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Load env failed in ora model")
		return nil, err
	}

	portStr := os.Getenv("port_oracle")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Error converting port to integer:", err)
		return nil, err
	}

	service_name := os.Getenv("service_name_oracle")
	hostname := os.Getenv("hostname_oracle")
	username := os.Getenv("username_oracle")
	password := os.Getenv("password_oracle")

	connStr := go_ora.BuildUrl(hostname, port, service_name, username, password, nil)
	conn, err := sql.Open("oracle", connStr)

	if err != nil {
		fmt.Println("error connecting to oracle database : ", err)
		return nil, err
	}

	fmt.Println("connection to oracle database success")
	return conn, nil
}
