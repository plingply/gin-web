package middleware

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

var ErrAbort = errors.New("手动停止")

func RecoverFunc(c *gin.Context) {
	fmt.Println("RecoverFunc")
	err := recover()
	fmt.Println("RecoverFunc => ", err)
	if err != nil {
		if err == ErrAbort {
			return
		}
	}
}

func DummyMiddleware(c *gin.Context) {
	defer RecoverFunc(c)
	fmt.Println("Im a dummy!")
	c.Next()
}

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := os.Getenv("Log_FILE_PATH")
	logFileName := os.Getenv("Log_FILE_NAME")

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	fmt.Println("log=>", fileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		fmt.Println("requres start")

		// 处理请求
		c.Next()

		fmt.Println("requres end")

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
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info("网络请求")
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("c.Request.Method==>", c.Request.Method)
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		// 处理请求
		c.Next()
	}
}
