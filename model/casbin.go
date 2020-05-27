package model

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"my_gin_cli/logger"
)

var Enforcer *casbin.Enforcer

// 初始化casbin
func CasbinSetup() {

	// 要使用自己定义的数据库rbac_db,最后的true很重要.默认为false,使用缺省的数据库名casbin,不存在则创建
	caAdapte, err := gormadapter.NewAdapter("mysql", viper.GetString("rbac.database"), true)
	if err != nil {
		logger.Logger.Debug("连接数据库错误: %v", err)
		return
	}

	caEnforcer, err := casbin.NewEnforcer("./conf/rbac_models.conf", caAdapte)
	if err != nil {
		logger.Logger.Debug("初始化casbin错误: %v", err)
		return
	}
	Enforcer = caEnforcer
}
