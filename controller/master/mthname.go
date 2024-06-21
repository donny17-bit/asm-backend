package master

import (
	"asm-backend/model"
	"fmt"
)

type MthNameData string

func Mthname() []MthNameData {
	fmt.Println("mthname controller")

	db, err := model.SqlModel()

	if err != nil {
		fmt.Println("error to connect to database")
		return nil
	}

	defer db.Close()

	query := "SELECT left(mthname, len(mthname) - 3) mthname FROM mthname order by thnblndesc"

	sql, err := db.Query(query)

	// check if there is error when execute the query
	if err != nil {
		fmt.Println("Error executing query row:", err)
		return nil
	}

	defer sql.Close()


	// Create an array to store the query results
	var mthnameData []MthNameData

	for sql.Next() {
		var data MthNameData

		// Scan each row into a struct
		err := sql.Scan(&data)
		if err != nil {
			fmt.Println("Error scanning row:", err)
		}

		// Append the struct to the array
		mthnameData = append(mthnameData, data)
	}

	if err := sql.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil
	}

	return mthnameData
}
