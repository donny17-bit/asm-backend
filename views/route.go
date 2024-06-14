package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Message string `json:"Message"`
}

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func GetDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}

func GetProductionLt(c *gin.Context) {
	c.HTML(http.StatusOK, "production.html", nil)
}

func GetProductionYr(c *gin.Context) {
	c.HTML(http.StatusOK, "production-yearly.html", nil)
}

