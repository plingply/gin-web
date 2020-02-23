package service

import (
	"fmt"
	"gin-web/middleware/jwt"
	"gin-web/models"
	"gin-web/repo"
	"gin-web/utils"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

type userServices struct {
}

func NewUserServices() *userServices {
	return &userServices{}
}

var userRepo = repo.NewUserRepository()

/*
登录
*/
func (u userServices) Login(username string, password string) (result models.Result) {

	if username == "" {
		result.Code = -1
		result.Msg = "请输入用户名！"
		return
	}
	if password == "" {
		result.Code = -1
		result.Msg = "请输入密码！"
		return
	}
	user := userRepo.GetUserByUserNameAndPwd(username, utils.GetMD5String(password))

	fmt.Println("user==>", user)

	if user.ID == 0 {
		result.Code = -1
		result.Msg = "用户名或密码错误!"
		return
	}
	// user.Token = middleware.GenerateToken(user)
	var claims jwt.CustomClaims
	claims.Username = user.Username
	claims.StandardClaims = jwtgo.StandardClaims{
		NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
		ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
		Issuer:    "newtrekWang",                   //签名的发行者
	}
	mjwt := jwt.NewJWT()
	token, err := mjwt.CreateToken(claims)
	if err != nil {
		result.Code = -1003
		result.Msg = err.Error()
		return
	}
	result.Code = 200
	result.Data = token
	result.Msg = "登录成功"
	return
}

// Add 添加
func (u userServices) Add(user models.User) (result models.Result) {
	code, p := userRepo.Save(user)
	if code == -1 {
		result.Code = -1
		result.Msg = "保存失败"
		return
	}
	result.Code = 200
	result.Msg = "创建成功"
	result.Data = p
	return
}

// Info 详情
func (u userServices) Info(user models.User) (result models.Result) {
	if user.Username == "" {
		result.Data = user.Username
		result.Code = -1
		result.Msg = "参数错误"
		return
	}
	agen := userRepo.GetUserByUsername(user.Username)
	if agen["username"] == "" {
		result.Data = nil
		result.Code = -1
		result.Msg = "用户不存在"
		return
	}
	result.Data = agen
	result.Code = 200
	result.Msg = "用户详情"
	return
}

// Delete 删除
func (u userServices) Delete(user models.User) (result models.Result) {
	if user.Username == "" {
		result.Data = user.Username
		result.Code = -1
		result.Msg = "参数错误"
		return
	}
	agen := userRepo.Deletes(user.Username)
	result.Data = agen
	result.Code = 200
	result.Msg = "删除成功"
	return
}

/*
保存
*/
func (u userServices) Save(user models.User) (result models.Result) {
	//添加
	if user.ID == 0 {
		agen := userRepo.GetUserByUsername(user.Username)
		if agen["id"] != 0 {
			result.Msg = "登录名重复,保存失败"
			return
		}
	}
	code, p := userRepo.Save(user)
	if code == -1 {
		result.Code = -1
		result.Msg = "保存失败"
		return
	}
	result.Code = 200
	result.Data = p
	return
}
