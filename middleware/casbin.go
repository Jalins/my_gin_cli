package middleware

import (
	"github.com/gin-gonic/gin"
	"my_gin_cli/logger"
	"my_gin_cli/model"
)

func Authorize() gin.HandlerFunc {

	return func(c *gin.Context) {

		var currentUser *model.User
		if user, ok := c.Get("user"); ok {
			if currentuser,ok :=  user.(*model.User); ok {
				currentUser = currentuser
			}
		}

		e := model.Enforcer
		logger.Logger.Debug("Enforcer", e)
		//从DB加载策略
		if err := e.LoadPolicy(); err != nil {
			logger.Logger.Debug("casbin加载策略失败！", err.Error())
		}

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := currentUser.Role
		logger.Logger.Debug("当前用户角色为：",currentUser.Role)

		//判断策略中是否存在
		if ok, _ := e.Enforce(sub, obj, act); ok {
			logger.Logger.Debug("恭喜您,权限验证通过")
			c.Next()
		} else {
			logger.Logger.Debug("很遗憾,权限验证没有通过")
			c.Abort()
		}
	}
}
