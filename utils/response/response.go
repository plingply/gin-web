package response

import (
	"gin-web/models"

	"github.com/gin-gonic/gin"
)

func Res(g *gin.Context, result models.Result) {
	g.JSON(200, result)
	return
}
