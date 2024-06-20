package views

import (
	"net/http"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Message string `json:"Message"`
}

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func GetDashboard(c *gin.Context) {
	// session := sessions.Default(c)
	// nama := session.Get("nama")

	// c.HTML(http.StatusOK, "dashboard.html", gin.H{
    //     "nama": nama,
    // })

	c.HTML(http.StatusOK, "dashboard.html", nil)
}

func GetProductionLt(c *gin.Context) {
	// session := sessions.Default(c)
	// nama := session.Get("nama")

	// ldc_id := session.Get("ldc_id") // default sesuai info logi
	// id := session.Get("id")
	// nik := session.Get("nik")
	// lastActivity := session.Get("lastActivity")
	// expiration := session.Get("expiration")

	// fmt.Println("route controller")
	// fmt.Println("session : ", session)
	// fmt.Println("id : ", id)
	// fmt.Println("nik : ", nik)
	// fmt.Println("ldc_id : ", ldc_id)
	// fmt.Println("lastActivity : ", lastActivity)
	// fmt.Println("expiration : ", expiration)


	c.HTML(http.StatusOK, "production.html", gin.H{
        "nama": "username",
    })
}

func GetProductionYr(c *gin.Context) {
	// session := sessions.Default(c)
	// nama := session.Get("nama")

	c.HTML(http.StatusOK, "production-yearly.html", gin.H{
        "nama": "username",
    })
}

