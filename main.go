package main

import (
	"my_gin_cli/conf"
	"my_gin_cli/router"

	_ "my_gin_cli/docs"
)

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
	r.Run(":8080")
}
