package views

import (
	"asm-backend/controller/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Define a struct to represent the JSON data
// type Data struct {
// 	Rn            string  `json:"Rn"`
// 	TglProd       string  `json:"TglProd"`
// 	ThnBln        string  `json:"ThnBln"`
// 	ProdDate      string  `json:"ProdDate"`
// 	BeginDate     string  `json:"BeginDate"`
// 	EndDate       string  `json:"EndDate"`
// 	Mo            string  `json:"Mo"`
// 	ClientName    string  `json:"ClientName"`
// 	Kanwil        string  `json:"Kanwil"`
// 	Cabang        string  `json:"Cabang"`
// 	Perwakilan    string  `json:"Perwakilan"`
// 	SubPerwakilan string  `json:"SubPerwakilan"`
// 	Jnner         string  `json:"Jnner"`
// 	JenisProd     string  `json:"JenisProd"`
// 	JenisPaket    *string `json:"JenisPaket"`
// 	Ket           *string `json:"Ket"`
// 	NamaCeding    *string `json:"NamaCeding"`
// 	Namaleader0   string  `json:"Namaleader0"`
// 	Namaleader1   string  `json:"Namaleader1"`
// 	Namaleader2   string  `json:"Namaleader2"`
// 	Namaleader3   string  `json:"Namaleader3"`
// 	GroupBusiness string  `json:"GroupBusiness"`
// 	Business      string  `json:"Business"`
// 	NoKontrak     *string `json:"NoKontrak"`
// 	NoPolis       string  `json:"NoPolis"`
// 	NoCif         string  `json:"NoCif"`
// 	ProdKe        string  `json:"ProdKe"`
// 	NamaDealer    *string `json:"NamaDealer"`
// 	Tsi           float32 `json:"Tsi"`
// 	Gpw           float32 `json:"Gpw"`
// 	Disc          float32 `json:"Disc"`
// 	Disc2         float32 `json:"Disc2"`
// 	Comm          float32 `json:"Comm"`
// 	Oc            float32 `json:"Oc"`
// 	Bkp           float32 `json:"Bkp"`
// 	Ngpw          float32 `json:"Ngpw"`
// 	Ri            float32 `json:"Ri"`
// 	Ricom         float32 `json:"Ricom"`
// 	Npw           float32 `json:"Npw"`
// }

type Data struct {
	Message string `json:"Message"`
}

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func GetDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}

func GetProduction(c *gin.Context) {

	c.HTML(http.StatusOK, "production.html", nil)
}

func GetProductionYearly(c *gin.Context) {
	c.HTML(http.StatusOK, "production-yearly.html", nil)
}

func GetProductionExample(c *gin.Context) {

	err := godotenv.Load()

	if err != nil {
		fmt.Print("Load env failed")
		return
	}

	auth_server := os.Getenv("auth_server")

	var ok bool

	if auth_server == "oracle" {
		ok = auth.IsActive(c)

		fmt.Println("is it okay : ", ok)
	}

	// if auth_server == "sql" {
	// 	ok = auth.IsActiveSql(c)
	// }

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

	session := sessions.Default(c)
	// ldc_id := session.Get("ldc_id") // default sesuai info login

	// if ldc_id == nil {
	// 	fmt.Println("error cabang kosong")
	// 	return
	// }

	id := session.Get("id")
	nik := session.Get("nik")
	lastActivity := session.Get("lastActivity")
	expiration := session.Get("expiration")

	fmt.Println("id : ", id)
	fmt.Println("nik : ", nik)
	fmt.Println("lastActivity : ", lastActivity)
	fmt.Println("expiration : ", expiration)

	resp, err := http.Get("http://localhost:8080/api/production-longterm?ldc_id=125&page=1&page_size=100&order=thnbln,%20client_name&begin_date=20240101&end_date=20240531")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching data from API")
		return
	}

	defer resp.Body.Close()

	// Parse the JSON response into a Data struct
	var data Data
	errJson := json.NewDecoder(resp.Body).Decode(&data)

	if errJson != nil {
		c.String(http.StatusInternalServerError, "Error decoding JSON data")
		return
	}

	fmt.Println("Data : ", data)
	fmt.Println("Data.Message : ", data.Message)
	// fmt.Println("Data.ProdDate : ", data.ProdDate)
	// fmt.Println("Data.ProdDate : ", data.ProdDate)
	// fmt.Println("Data.JenisProd : ", data.JenisProd)

	// Render the HTML template with the data
	c.HTML(http.StatusOK, "production.html", gin.H{
		// "ProdDate":  data.ProdDate,
		// "JenisProd": data.JenisProd,
		"Message": data.Message,
	})
}
