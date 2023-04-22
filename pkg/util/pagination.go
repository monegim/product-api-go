package util

import (
	"github.com/gin-gonic/gin"
	"github.com/monegim/product-api-go/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	pageStr := c.Query("page")
	page := MustInt(pageStr)
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}
