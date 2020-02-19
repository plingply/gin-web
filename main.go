package main

import (
	"gin-web/middleware"
	router "gin-web/router"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := gin.Default()
	// 全局中间件
	// Logger 中间件将写日志到 gin.DefaultWriter ,即使你设置 GIN_MODE=release 。
	// 默认 gin.DefaultWriter = os.Stdout

	server.Use(middleware.LoggerToFile())

	server.Use(middleware.DummyMiddleware)

	// Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误。
	server.Use(gin.Recovery())

	server.Static("/static", "./static")

	router.InitRouter(server)

	port := os.Getenv("PORT")

	server.Run(":" + port)

}
