package main

import (
	"github.com/gin-gonic/gin"
	"github.com/monegim/product-api-go/pkg/setting"
)

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
}

func GetCoffeeHandler(c *gin.Context) {

}
