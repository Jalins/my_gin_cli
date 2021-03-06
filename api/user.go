package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"my_gin_cli/controller"
	"my_gin_cli/serializer"
)

// @Summary 用户注册
// @Tags 用户
// @version 1.0
// @Accept application/json
// @Param UserRegisterService body service.UserRegisterService true "往数据库中插入一个新的用户"
// @Success 200 object serializer.Response 成功后返回值
// @Failure 500 object serializer.Response 注册失败
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	var service controller.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// @Summary 用户登录
// @Tags 用户
// @version 1.0
// @Accept application/json
// @Param UserLoginService body service.UserLoginService true "验证该用户的账户是否在数据库存在"
// @Success 200 object serializer.Response 成功后返回值
// @Failure 500 object serializer.Response 登录失败
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var service controller.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// @Summary 用户列表
// @Tags 用户
// @version 1.0
// @Accept application/json
// @Param UserListService body service.UserListService true "查询所有的用户"
// @Success 200 object serializer.Response 成功后返回值
// @Failure 500 object serializer.Response 登录失败
// @Router /user/list [get]
func UserList(c *gin.Context) {
	var service controller.UserLoginService
	if res, err := service.List(); err == nil {
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// @Summary 用户详情
// @Tags 用户
// @version 1.0
// @Accept application/json
// @Success 200 object serializer.Response 成功后返回值
// @Failure 500 object serializer.Response 查询失败
// @Router /user/me [get]
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.Response{
		Code:  200,
		Data:  *user,
	}
	c.JSON(200, res)
}

// @Summary 用户退出
// @Tags 用户
// @version 1.0
// @Accept application/json
// @Success 200 object serializer.Response 成功后返回值
// @Failure 500 object serializer.Response 退出失败
// @Router /user/logout [get]
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "退出成功",
	})
}
