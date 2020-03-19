package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"my_gin_cli/util"
	"os"
	"path"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := viper.GetString("logger.log_file_path")
	logFileName := viper.GetString("logger.log_file_name")

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	// 判断文件是否存在
	if _, bool := util.IsFileExist(fileName); bool == false {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println("创建文件失败")
		}
	}

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("打开文件错误：", err)
	}

	// 实例化
	logger := logrus.New()

	// 设置输出
	// 输出到文件中
	logger.Out = src
	// 输出到控制台
	logger.Out = os.Stdout

	// 设置日志级别
	if viper.GetString("logger.log_level") == "trace" {
		logger.SetLevel(logrus.TraceLevel)
	}else if viper.GetString("logger.log_level") == "debug" {
		logger.SetLevel(logrus.DebugLevel)
	}else if viper.GetString("logger.log_level") == "info" {
		logger.SetLevel(logrus.InfoLevel)
	}else if viper.GetString("logger.log_level") == "warning" {
		logger.SetLevel(logrus.WarnLevel)
	}else if viper.GetString("logger.log_level") == "error" {
		logger.SetLevel(logrus.ErrorLevel)
	}else if viper.GetString("logger.log_level") == "fatal" {
		logger.SetLevel(logrus.FatalLevel)
	}else if viper.GetString("logger.log_level") == "panic" {
		logger.SetLevel(logrus.PanicLevel)
	}else {
		fmt.Printf("not a valid logrus level")
	}


	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName + ".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code"  : statusCode,
			"latency_time" : latencyTime,
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
		}).Info()
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
