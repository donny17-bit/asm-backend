package web

import (
	"asm-backend/auth"
	"fmt"
	"net/http"

	"asm-backend/model"

	"github.com/gin-gonic/gin"
)

func GetBusinessSource(c *gin.Context) {

	ok := auth.IsActive(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"data":           "",
			"current_page: ": "",
			"page_size":      "",
			"max_page":       "",
			"message":        "unauthorized",
			"status":         401,
		})
		return
	}

	lvl := c.Query("lvl")
	nama := c.Query("nama")
	// tambahin filter by namaleader biar datanya ga banyak

	db, err := model.SqlModel()

	if err != nil {
		fmt.Println("error to connect to database")
		return
	}
	defer db.Close()

	// get query
	var queryFinal string
	query := "select Lag_agen_id, lvl, lag_leader, nama, leader0, leader1, leader2, leader3, " +
		"leader4, leader5, namaleader0, namaleader1, namaleader2, namaleader3, namaleader4, namaleader5 " +
		"from mv_agen where lvl = '" + lvl + "' "

	if nama != "" {
		queryFinal = query + " and nama like '%" + nama + "%' order by nama"
	} else {
		queryFinal = query + " order by nama"
	}

	rows, err := db.Query(queryFinal)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close() // Close the result set when done

	type Data struct {
		LagAgenId   string
		Lvl         string
		LagLeader   *string
		Nama        string
		Leader0     string
		Leader1     string
		Leader2     string
		Leader3     string
		Leader4     string
		Leader5     string
		Namaleader0 string
		Namaleader1 string
		Namaleader2 string
		Namaleader3 string
		Namaleader4 string
		Namaleader5 string
	}

	var datas []Data

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.LagAgenId, &data.Lvl, &data.LagLeader, &data.Nama, &data.Leader0, &data.Leader1, &data.Leader2,
			&data.Leader3, &data.Leader4, &data.Leader5, &data.Namaleader0, &data.Namaleader1, &data.Namaleader2, &data.Namaleader3,
			&data.Namaleader4, &data.Namaleader5)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		// Append the struct to the array
		datas = append(datas, data)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    datas,
		"message": "success get data",
	})
}
