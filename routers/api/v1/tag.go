package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/monegim/product-api-go/pkg/app"
	"github.com/monegim/product-api-go/pkg/setting"
	"github.com/monegim/product-api-go/pkg/util"
	"github.com/monegim/product-api-go/service/tag_service"
)

// @Summary Get multiple article tags
func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = util.MustInt(arg)
	}
	tagService := tag_service.Tag{
		Name: name,
		State: state,
		PageSize: setting.AppSetting.PageSize,
		PageNum: util.GetPage(c),
	}
	
}
