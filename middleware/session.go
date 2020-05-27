package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"my_gin_cli/logger"
)

// Session 初始化session
func Session() gin.HandlerFunc {
	store, err := redis.NewStore(50, "tcp", viper.GetString("redis.redis_addr"), "", []byte(viper.GetString("session.session_secret")))
	if err != nil {
		logger.Logger.Error("redis.NewStore出错",err.Error())
	}

	return sessions.Sessions("mysession", store)
}
