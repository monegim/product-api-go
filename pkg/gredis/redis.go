package gredis

import (
	"context"
	"net"

	"github.com/monegim/product-api-go/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var RedisConn *redis.Client

func Setup() error {
	RedisConn = redis.NewClient(&redis.Options{
		MaxIdleConns:    setting.RedisSetting.MaxIdle,
		ConnMaxIdleTime: setting.RedisSetting.IdleTimeout,
		Addr:            setting.RedisSetting.Host,
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			c, err :=  net.Dial(network, addr)
			if err != nil {
				return nil, err
			}
			// TODO: Handle password
			return c, err	
		},
	})
	return nil
}
