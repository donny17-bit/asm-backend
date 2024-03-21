package web

import (
	"asm-backend/auth"
	"asm-backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductionYr(c *gin.Context) {

	ok := auth.IsActive(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
			"status":  401,
		})
		return
	}

	no_polis := c.Query("no_polis")

	fmt.Println("no_polis : ", no_polis)

	db, err := model.SqlModel()

	if err != nil {
		fmt.Println("error to connect to database")
		return
	}
	defer db.Close()

	// Perform a query
	query := "exec Warehouse_Asm.dbo.SP_PRINT_DETAIL_PRODUCTION_LONGTERM "
	where := "'where no_polis = ''" + no_polis + "'' '" // nnti diganti pake where in

	rows, err := db.Query(query + where)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close() // Close the result set when done

	type Data struct {
		ThnBln        string  `json:"ThnBln"`
		ProdDate      string  `json:"ProdDate"`
		BeginDate     string  `json:"BeginDate"`
		EndDate       string  `json:"EndDate"`
		Mo            string  `json:"Mo"`
		ClientName    string  `json:"ClientName"`
		Kanwil        string  `json:"Kanwil"`
		Cabang        string  `json:"Cabang"`
		Perwakilan    string  `json:"Perwakilan"`
		SubPerwakilan string  `json:"SubPerwakilan"`
		Jnner         string  `json:"Jnner"`
		JenisProd     string  `json:"JenisProd"`
		JenisPaket    *string `json:"JenisPaket"`
		Ket           string  `json:"Ket"`
		NamaCeding    *string `json:"NamaCeding"`
		Namaleader0   string  `json:"Namaleader0"`
		Namaleader1   string  `json:"Namaleader1"`
		Namaleader2   string  `json:"Namaleader2"`
		Namaleader3   string  `json:"Namaleader3"`
		GroupBusiness string  `json:"GroupBusiness"`
		Business      string  `json:"Business"`
		NoKontrak     *string `json:"NoKontrak"`
		Alasan        *string `json:"Alasan"`
		NoPolis       string  `json:"NoPolis"`
		NoCif         string  `json:"NoCif"`
		ProdKe        string  `json:"ProdKe"`
		NamaDealer    *string `json:"NamaDealer"`
		Tsi           float32 `json:"Tsi"`
		Gpw           float32 `json:"Gpw"`
		Disc          float32 `json:"Disc"`
		Disc2         float32 `json:"Disc2"`
		Comm          float32 `json:"Comm"`
		Oc            float32 `json:"Oc"`
		Bkp           float32 `json:"Bkp"`
		Ngpw          float32 `json:"Ngpw"`
		Ri            float32 `json:"Ri"`
		Ricom         float32 `json:"Ricom"`
		Npw           float32 `json:"Npw"`
	}

	// Create an array to store the query results
	var datas []Data

	for rows.Next() {
		var data Data

		// Scan each row into a struct
		if err := rows.Scan(&data.ThnBln, &data.ProdDate, &data.BeginDate,
			&data.EndDate,
			&data.Mo,
			&data.ClientName,
			&data.Kanwil,
			&data.Cabang,
			&data.Perwakilan,
			&data.SubPerwakilan,
			&data.Jnner,
			&data.JenisProd,
			&data.JenisPaket,
			&data.Ket,
			&data.NamaCeding,
			&data.Namaleader0,
			&data.Namaleader1,
			&data.Namaleader2,
			&data.Namaleader3,
			&data.GroupBusiness,
			&data.Business,
			&data.NoKontrak,
			&data.Alasan,
			&data.NoPolis,
			&data.NoCif,
			&data.ProdKe,
			&data.NamaDealer,
			&data.Tsi,
			&data.Gpw,
			&data.Disc,
			&data.Disc2,
			&data.Comm,
			&data.Oc,
			&data.Bkp,
			&data.Ngpw,
			&data.Ri,
			&data.Ricom,
			&data.Npw,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Append the struct to the array
		datas = append(datas, data)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}

	fmt.Print("data : ", datas)

	// if no error
	// Respond with JSON data
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    datas,
		"message": "success get data",
	})
}
