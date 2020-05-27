package router

import (
	"github.com/spf13/viper"
	"my_gin_cli/api"
	"my_gin_cli/middleware"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 使用swagger文档
	url := ginSwagger.URL(viper.GetString("swagger.swagger_addr"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 中间件, 顺序不能改
	r.Use(middleware.Session())
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		// 用户注册
		v1.POST("user/register", api.UserRegister)
		// 用户登录
		v1.POST("user/login", api.UserLogin)

		admin := v1.Group("")
		{
			admin.Use(middleware.Authorize())
			// 需要admin用户才能操作
			admin.GET("admin/remove", api.Remove)
		}


		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
