package gredis

import (
	"github.com/monegim/product-api-go/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var RedisConn *redis.Client

func Setup() error {
	RedisConn := redis.NewClient(&redis.Options{
		MaxIdleConns:    setting.RedisSetting.MaxIdle,
		ConnMaxIdleTime: setting.RedisSetting.IdleTimeout,
		Addr:            setting.RedisSetting.Host,
	})

}
