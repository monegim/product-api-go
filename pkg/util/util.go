package util

import (
	"strconv"

	"github.com/monegim/product-api-go/pkg/setting"
)

func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

func MustInt(s string) int {
	v, _ := strconv.ParseInt(s, 10, 0)
	return int(v)
}
