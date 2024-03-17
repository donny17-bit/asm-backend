package model

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	go_ora "github.com/sijms/go-ora/v2"
)

func OraModel() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Load env failed in ora model")
		return nil
	}

	portStr := os.Getenv("port_oracle")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Error converting port to integer:", err)
		return nil
	}

	service_name := os.Getenv("service_name_oracle")
	hostname := os.Getenv("hostname_oracle")
	username := os.Getenv("username_oracle")
	password := os.Getenv("password_oracle")

	connStr := go_ora.BuildUrl(hostname, port, service_name, username, password, nil)
	conn, err := sql.Open("oracle", connStr)

	if err != nil {
		fmt.Println("error connecting to oracle database : ", err)
		return nil
	}

	fmt.Println("success connect to oracle database")
	return conn
}
