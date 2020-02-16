package base_controller

import (
	"fmt"
	"gin-web/middleware/jwt"
	"gin-web/models"
	"gin-web/utils/response"

	"github.com/gin-gonic/gin"
)

func RefreshToken(g *gin.Context) {
	var result models.Result
	tokenParams := g.PostForm("token")
	if tokenParams == "" {
		result.Code = -1
		result.Msg = "参数错误"
		result.Data = nil
		response.Res(g, result)
		return
	}

	fmt.Println("tokenParams", tokenParams)
	mjwt := jwt.NewJWT()
	token, err := mjwt.RefreshToken(tokenParams)
	if err != nil {
		result.Code = -1000
		result.Msg = err.Error()
		result.Data = nil
		response.Res(g, result)
		return
	}

	result.Code = 200
	result.Msg = ""
	result.Data = token
	response.Res(g, result)
}
