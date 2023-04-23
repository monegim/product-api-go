package gredis

import (
	"context"
	"encoding/json"
	"net"
	"time"

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

// Set a key/value
func Set(key string, data interface{}, expiration int) error {
	conn := RedisConn.Conn()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	statusCmd := conn.Set(context.TODO(), key, value, time.Duration(expiration)*time.Second)
	return statusCmd.Err()
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
