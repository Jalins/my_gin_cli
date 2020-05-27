package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"my_gin_cli/logger"
	"my_gin_cli/model"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	viper.SetConfigName("config") 				// 设置配置文件名 (不带后缀)
	viper.AddConfigPath("./conf")    			// 第一个搜索路径
	viper.SetConfigType("json")					// 配置文件的格式
	if err := viper.ReadInConfig(); err != nil { 	// 读取配置数据
		logger.Logger.Debug("读取配置文件失败！", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Logger.Debug("配置发生变更：", e.Name)
	})


	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		logger.Logger.Debug("翻译文件加载失败", err)
	}

	// 加载casbin
	model.CasbinSetup()

	// 连接数据库
	model.Database(viper.GetString("database.mysql"))

	// 开启debug模式
	gin.SetMode(gin.DebugMode)
	// 开始release模式
	//gin.SetMode(gin.ReleaseMode)
}
