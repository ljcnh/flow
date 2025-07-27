package red

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	client *redis.Client
)

// InitRedis 初始化Redis连接
func InitRedis() error {
	var (
		err error
	)
	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		return fmt.Errorf("未设置redis密码")
	}
	client = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: password,
		DB:       viper.GetInt("redis.db"),

		// 连接池配置
		PoolSize:        viper.GetInt("redis.pool_size"),
		MinIdleConns:    viper.GetInt("redis.min_idle_conns"),
		ConnMaxIdleTime: 30 * time.Second,

		// 超时设置
		DialTimeout:  viper.GetDuration("redis.dial_timeout"),
		ReadTimeout:  viper.GetDuration("redis.read_timeout"),
		WriteTimeout: viper.GetDuration("redis.write_timeout"),

		// 网络配置
		MaxRetries:      3,
		MinRetryBackoff: 100 * time.Millisecond,
		MaxRetryBackoff: 1 * time.Second,
	})
	if err = client.Ping(context.Background()).Err(); err != nil {
		return fmt.Errorf("redis连接失败: %w", err)
	}
	return nil
}

func GetClient() *redis.Client {
	return client
}
