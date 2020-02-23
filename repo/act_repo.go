package repo

import (
	"fmt"
	"gin-web/datasource"
	"gin-web/models"
	"log"
)

type ActRepository interface {
	GetList(page int, limit int) map[string]interface{}
	Save(act models.Article) (int, models.Article)
	GetActicleInfo(id uint) models.Params
	Deletes(id uint) models.Article
}

type actRepositorys struct{}

func NewActRepository() ActRepository {
	return &actRepositorys{}
}

//获取列表
func (n actRepositorys) GetList(page int, limit int) (params map[string]interface{}) {

	var count int = 0
	db := datasource.GetDB()

	rows, err := db.Table("article").Offset(page*limit - limit).Select("article.*, type.name as type_name").Joins("left join type on type.code = article.type").Limit(limit).Count(&count).Rows()

	if err != nil {
		return nil
	}

	param, errs := models.FormatResult(rows)

	if errs != nil {
		return nil
	}

	params = make(map[string]interface{})
	params["count"] = count
	params["item"] = param
	params["limit"] = limit
	params["page"] = page

	return
}

//添加/修改
func (n actRepositorys) Save(act models.Article) (int, models.Article) {
	code := 0
	db := datasource.GetDB()
	if act.ID != 0 {
		if err := db.Save(&act).Error; err != nil {
			log.Println("报错了==>", err)
			code = -1
		}
		return code, act
	}
	if err := db.Create(&act).Error; err != nil {
		fmt.Println("报错了==>", err)
		code = -1
	}
	return code, act
}

func (n actRepositorys) GetActicleInfo(id uint) (param models.Params) {
	db := datasource.GetDB()
	rows, err := db.Table("article").Select("article.*, type.name as type_name, type.code as type_code").Joins("left join type on type.code = article.type").Where("article.id = ?", id).Rows()

	if err != nil {
		return nil
	}

	params, errs := models.FormatResult(rows)

	fmt.Println("params", params)

	if errs != nil {
		return nil
	}

	param = params[0]

	return
}

func (n actRepositorys) Deletes(id uint) models.Article {
	var act models.Article
	db := datasource.GetDB()
	db.Where("id = ?", id).Delete(&act)
	return act
}
