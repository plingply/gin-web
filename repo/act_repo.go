package repo

import (
	"fmt"
	"gin-web/datasource"
	"gin-web/models"
)

type ActRepository interface {
	GetList(page int, limit int) map[string]interface{}
}

func NewActRepository() ActRepository {
	return &actRepository{}
}

type actRepository struct{}

//获取列表
func (n actRepository) GetList(page int, limit int) (params map[string]interface{}) {

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
