package api

import (
	"github.com/gin-gonic/gin"
	"github.com/monegim/product-api-go/pkg/app"
)

type auth struct {
	Username string `json:"username" binding:"Required"`
	Password string `json:"password" binding:"Required"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.PostForm("username")
	password := c.PostForm("password")
	
	a := auth{Username: username, Password: password}
	// ok, _ := 
}
