package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"my_gin_cli/cache"
	"my_gin_cli/model"
	"my_gin_cli/util"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	viper.SetConfigName("config") 				// 设置配置文件名 (不带后缀)
	viper.AddConfigPath("./conf")    			// 第一个搜索路径
	viper.SetConfigType("json")					// 配置文件的格式
	if err := viper.ReadInConfig(); err != nil { 	// 读取配置数据
		fmt.Println("读取配置文件失败！", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变更：", e.Name)
	})


	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(viper.GetString("database.mysql"))
	cache.Redis()

	// 开启debug模式
	gin.SetMode(gin.DebugMode)
	// 开始release模式
	//gin.SetMode(gin.ReleaseMode)
}
