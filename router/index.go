package router

import (
	"gin-web/controllers/act_controller"
	"gin-web/controllers/base_controller"
	"gin-web/controllers/file_controller"
	"gin-web/controllers/user_controller"
	"gin-web/middleware/jwt"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(router *gin.Engine) {

	router.LoadHTMLGlob("static/*.html")

	// 模板
	view := router.Group("/")
	{
		view.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", nil)
		})

		view.GET("/upload", func(c *gin.Context) {
			c.HTML(200, "upload.html", nil)
		})
	}

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

	// 文章
	act := api.Group("/article")
	{
		act.GET("/list", act_controller.List)
	}

	// 文件上传
	file := api.Group("/file")
	{
		file.POST("/upload", file_controller.UpLoad)
	}
}
