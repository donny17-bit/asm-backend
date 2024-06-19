package views

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Message string `json:"Message"`
}

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func GetDashboard(c *gin.Context) {
	session := sessions.Default(c)
	nama := session.Get("nama")

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
        "nama": nama,
    })
}

func GetProductionLt(c *gin.Context) {
	session := sessions.Default(c)
	nama := session.Get("nama")

	c.HTML(http.StatusOK, "production.html", gin.H{
        "nama": nama,
    })
}

func GetProductionYr(c *gin.Context) {
	session := sessions.Default(c)
	nama := session.Get("nama")

	c.HTML(http.StatusOK, "production-yearly.html", gin.H{
        "nama": nama,
    })
}

