package cache

import (
	"github.com/spf13/viper"
	"my_gin_cli/util"
	"strconv"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(viper.GetString("redis.redis_db"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       viper.GetString("redis.redis_addr"),
		Password:   viper.GetString("redis.redis_pw"),
		DB:         int(db),
		MaxRetries: 1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		util.Log().Panic("连接Redis不成功", err)
	}

	RedisClient = client
}
