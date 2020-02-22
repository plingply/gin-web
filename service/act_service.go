package service

import (
	"gin-web/models"
	"gin-web/repo"
)

type ActService interface {
	List(page int, limit int) (result models.Result)
}

type actServices struct {
}

func NewActServices() *actServices {
	return &actServices{}
}

var actRepo = repo.NewActRepository()

// List 列表
func (u actServices) List(page int, limit int) (result models.Result) {

	if page == 0 || limit == 0 {
		result.Data = nil
		result.Code = -1
		result.Msg = "参数错误"
		return
	}

	agen := actRepo.GetList(page, limit)

	result.Data = agen
	result.Code = 200
	result.Msg = "文章列表"
	return
}
