package gredis

import (
	"context"
	"net"

	"github.com/monegim/product-api-go/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var RedisConn *redis.Client

func Setup() error {
	logrus.Info("Initializing gredis")
	RedisConn = redis.NewClient(&redis.Options{
		MaxIdleConns:    setting.RedisSetting.MaxIdle,
		ConnMaxIdleTime: setting.RedisSetting.IdleTimeout,
		Addr:            setting.RedisSetting.Host,
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			c, err := net.Dial(network, addr)
			if err != nil {
				return nil, err
			}
			// TODO: Handle password
			return c, err
		},
	})
	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Conn()
	defer conn.Close()
	// conn.Exist
	intCmd := conn.Exists(context.TODO(), key)
	return intCmd.Val() != 0
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Conn()
	defer conn.Close()
	strCmd := RedisConn.Get(context.TODO(), key)
	return strCmd.Bytes()
}
