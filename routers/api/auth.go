package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/monegim/product-api-go/pkg/app"
	"github.com/monegim/product-api-go/pkg/e"
	"github.com/monegim/product-api-go/pkg/util"
	"github.com/monegim/product-api-go/service/auth_service"
)

type auth struct {
	Username string `json:"username" binding:"Required"`
	Password string `json:"password" binding:"Required"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.PostForm("username")
	password := c.PostForm("password")

	// a := auth{Username: username, Password: password}
	//TODO
	// validate user, pass
	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}
	token, err := util.GenerateToken()
}
