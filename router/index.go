package router

import (
	"gin-web/controllers/base_controller"
	"gin-web/controllers/user_controller"
	"gin-web/middleware/jwt"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(router *gin.Engine) {
	api := router.Group("/api")
	api.POST("/login", user_controller.Login)
	api.POST("/refresh/token", base_controller.RefreshToken)
	// 用户
	user := api.Group("/user")
	{
		user.Use(jwt.JWTAuth())
		user.POST("/add", user_controller.Add)
		user.POST("/delete/:username", user_controller.Delete)
		user.GET("/info/:username", user_controller.Info)
	}

}
