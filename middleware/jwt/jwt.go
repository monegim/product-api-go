package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/monegim/product-api-go/pkg/e"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
	}
}
