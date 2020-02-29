package user_controller

import (
	"fmt"
	"gin-web/models"
	"gin-web/service"
	"gin-web/utils/response"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

// UserService 服务
var UserService = service.NewUserServices()

// Login 登录
func Login(g *gin.Context) {
	username := g.PostForm("username")
	password := g.PostForm("password")

	fmt.Println(username, password)
	result := UserService.Login(username, password)
	response.Res(g, result)
}

// Add 保存
func Add(g *gin.Context) {
	var user models.User
	// 绑定参数到结构体
	username := g.PostForm("username")
	password := g.PostForm("password")
	name := g.PostForm("name")
	mobile := g.PostForm("mobile")
	email := g.PostForm("email")
	user.Username = username
	user.Password = password
	user.Name = name
	user.Mobile = mobile
	user.Email = email
	fmt.Println("Bind", user)
	result := UserService.Add(user)
	response.Res(g, result)
}

// Save 保存
func Save(g *gin.Context) {
	var user models.User
	result := UserService.Save(user)
	response.Res(g, result)
}

// Delete 删除
func Delete(g *gin.Context) {
	var user models.User
	username := g.Param("username")
	user.Username = username
	result := UserService.Delete(user)
	response.Res(g, result)
}

// Info 详情
func Info(g *gin.Context) {
	if token, isOk := g.Get("claims"); isOk {
		fmt.Println("tokne.user=>>", token)
	}

	var user models.User
	username := g.Param("username")
	user.Username = username
	result := UserService.Info(user)
	response.Res(g, result)
}

// InfoByID 详情
func InfoByID(g *gin.Context) {
	if token, isOk := g.Get("claims"); isOk {
		fmt.Println("tokne.user=>>", token)
	}

	var user models.User
	id := g.Param("id")
	ids := cast.ToUint(id)
	user.ID = ids
	result := UserService.InfoByID(user)
	response.Res(g, result)
}
