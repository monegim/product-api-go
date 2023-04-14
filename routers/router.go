package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/monegim/product-api-go/routers/api"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.)
	return r
}
