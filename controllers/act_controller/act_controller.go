package act_controller

import (
	"gin-web/service"
	"gin-web/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// ActService 服务
var ActService = service.NewActServices()

func List(g *gin.Context) {

	page := cast.ToInt(g.Query("page"))
	limit := cast.ToInt(g.Query("limit"))

	result := ActService.List(page, limit)
	response.Res(g, result)
}
