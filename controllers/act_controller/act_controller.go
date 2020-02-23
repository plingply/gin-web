package act_controller

import (
	"gin-web/models"
	"gin-web/service"
	"gin-web/utils/response"
	"time"

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

func Add(g *gin.Context) {

	var act models.Article
	Title := g.PostForm("title")
	Content := g.PostForm("content")
	Remark := g.PostForm("remark")
	VideoURL := g.PostForm("video_url")
	VideoCover := g.PostForm("video_cover")
	Picture := g.PostForm("picture")
	Type := g.PostForm("type")

	var CreatedAt = time.Now()
	var UpdatedAt = time.Now()

	act.Title = Title
	act.Content = Content
	act.Remark = Remark
	act.VideoURL = VideoURL
	act.VideoCover = VideoCover
	act.Picture = Picture
	act.Type = Type
	act.CreatedAt = CreatedAt
	act.UpdatedAt = UpdatedAt

	result := ActService.Add(act)
	response.Res(g, result)
}

func Update(g *gin.Context) {

	var act models.Article
	Title := g.PostForm("title")
	Content := g.PostForm("content")
	Remark := g.PostForm("remark")
	VideoURL := g.PostForm("video_url")
	VideoCover := g.PostForm("video_cover")
	Picture := g.PostForm("picture")
	Type := g.PostForm("type")
	ID := g.PostForm("id")

	ids := cast.ToUint(ID)

	var CreatedAt = time.Now()
	var UpdatedAt = time.Now()

	act.ID = ids
	act.Title = Title
	act.Content = Content
	act.Remark = Remark
	act.VideoURL = VideoURL
	act.VideoCover = VideoCover
	act.Picture = Picture
	act.Type = Type
	act.CreatedAt = CreatedAt
	act.UpdatedAt = UpdatedAt

	result := ActService.Update(act)
	response.Res(g, result)
}

// Delete 删除
func Delete(g *gin.Context) {
	id := g.Param("id")
	ids := cast.ToUint(id)
	result := ActService.Delete(ids)
	response.Res(g, result)
}

// Info 详情
func Info(g *gin.Context) {
	id := g.Param("id")
	ids := cast.ToUint(id)
	result := ActService.Info(ids)
	response.Res(g, result)
}
