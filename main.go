package main

import (
	"fmt"
	"my_gin_cli/conf"
	"my_gin_cli/logger"
	"my_gin_cli/router"

	_ "my_gin_cli/docs"
)

func init() {
	logger.LogConf()
}

// @title 项目名称 swagger
// @version 1.0
// @description 项目名称 swagger api文档
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := router.NewRouter()


	if err := r.Run(":8080"); err != nil {
		logger.Logger.Debug("服务器启动失败！", err)
		fmt.Println("服务器启动失败！",err)
	}
}
