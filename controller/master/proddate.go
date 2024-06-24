package master

import (
	"asm-backend/model"
	"fmt"
)

type ProdDateTime string

func Proddate() []ProdDateTime {
	fmt.Println("proddate controller")

	db, err := model.SqlModel()

	if err != nil {
		fmt.Println("error to connect to database")
		return nil
	}

	defer db.Close()

	query := "select proddatetime2 as proddatetime from proddatetime where thnblndesc >= '0' order by thnblndesc"

	sql, err := db.Query(query)

	// check if there is error when execute the query
	if err != nil {
		fmt.Println("Error executing query row:", err)
		return nil
	}

	defer sql.Close()


	// Create an array to store the query results
	var proddateData []ProdDateTime

	for sql.Next() {
		var data ProdDateTime

		// Scan each row into a struct
		err := sql.Scan(&data)
		if err != nil {
			fmt.Println("Error scanning row:", err)
		}

		// Append the struct to the array
		proddateData = append(proddateData, data)
	}

	if err := sql.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil
	}

	return proddateData
}
