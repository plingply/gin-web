package service

import (
	"gin-web/models"
	"gin-web/repo"
)

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

func (u actServices) Add(act models.Article) (result models.Result) {

	if act.Title == "" {
		result.Code = -1
		result.Msg = "标题不能为空"
		return
	}

	if act.Content == "" {
		result.Code = -1
		result.Msg = "内容不能为空"
		return
	}

	if act.Type == "" {
		result.Code = -1
		result.Msg = "类型不能为空"
		return
	}

	code, p := actRepo.Save(act)
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

func (u actServices) Update(act models.Article) (result models.Result) {

	if act.ID == 0 {
		result.Code = -1
		result.Msg = "id不能为空"
		return
	}

	if act.Title == "" {
		result.Code = -1
		result.Msg = "标题不能为空"
		return
	}

	if act.Content == "" {
		result.Code = -1
		result.Msg = "内容不能为空"
		return
	}

	if act.Type == "" {
		result.Code = -1
		result.Msg = "类型不能为空"
		return
	}

	code, p := actRepo.Save(act)
	if code == -1 {
		result.Code = -1
		result.Msg = "保存失败"
		return
	}
	result.Code = 200
	result.Msg = "保存成功"
	result.Data = p
	return
}

// Info 详情
func (u actServices) Info(id uint) (result models.Result) {
	if id == 0 {
		result.Data = nil
		result.Code = -1
		result.Msg = "参数错误"
		return
	}
	agen := actRepo.GetActicleInfo(id)

	if agen == nil {
		result.Data = nil
		result.Code = -1
		result.Msg = "获取用户详情失败"
		return
	}
	result.Data = agen
	result.Code = 200
	result.Msg = "用户详情"
	return
}

// Delete 删除
func (u actServices) Delete(id uint) (result models.Result) {
	if id == 0 {
		result.Data = nil
		result.Code = -1
		result.Msg = "参数错误"
		return
	}
	agen := actRepo.Deletes(id)
	result.Data = agen
	result.Code = 200
	result.Msg = "删除成功"
	return
}
