package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

//实例化
var logger = logrus.New()

func log(c *gin.Context) map[string]interface{} {

	logFilePath := os.Getenv("Log_FILE_PATH")
	logFileName := os.Getenv("Log_FILE_NAME")

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.ErrorLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var result = make(map[string]interface{})
	// 日志格式
	if c == nil {
		result["status_code"] = ""
		result["latency_time"] = ""
		result["client_ip"] = ""
		result["req_method"] = ""
		result["req_uri"] = ""
	} else {
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
		result["status_code"] = statusCode
		result["latency_time"] = latencyTime
		result["client_ip"] = clientIP
		result["req_method"] = reqMethod
		result["req_uri"] = reqUri
	}

	return result
}

func Fatal(msg string, c *gin.Context) {

	result := log(c)

	logger.WithFields(logrus.Fields{
		"status_code":  result["status_code"],
		"latency_time": result["latency_time"],
		"client_ip":    result["client_ip"],
		"req_method":   result["req_method"],
		"req_uri":      result["req_uri"],
	}).Fatal(msg)
}

func Error(msg string, c *gin.Context) {

	result := log(c)

	logger.WithFields(logrus.Fields{
		"status_code":  result["status_code"],
		"latency_time": result["latency_time"],
		"client_ip":    result["client_ip"],
		"req_method":   result["req_method"],
		"req_uri":      result["req_uri"],
	}).Error(msg)
}

func Warn(msg string, c *gin.Context) {

	result := log(c)

	logger.WithFields(logrus.Fields{
		"status_code":  result["status_code"],
		"latency_time": result["latency_time"],
		"client_ip":    result["client_ip"],
		"req_method":   result["req_method"],
		"req_uri":      result["req_uri"],
	}).Warn(msg)
}

func Info(msg string, c *gin.Context) {

	result := log(c)

	logger.WithFields(logrus.Fields{
		"status_code":  result["status_code"],
		"latency_time": result["latency_time"],
		"client_ip":    result["client_ip"],
		"req_method":   result["req_method"],
		"req_uri":      result["req_uri"],
	}).Info(msg)
}

func Debug(msg string, c *gin.Context) {

	result := log(c)

	logger.WithFields(logrus.Fields{
		"status_code":  result["status_code"],
		"latency_time": result["latency_time"],
		"client_ip":    result["client_ip"],
		"req_method":   result["req_method"],
		"req_uri":      result["req_uri"],
	}).Debug(msg)
}
