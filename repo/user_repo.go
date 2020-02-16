package repo

import (
	"gin-web/datasource"
	"gin-web/models"
	"log"

	"gin-web/utils"
)

type UserRepository interface {
	GetUserByUserNameAndPwd(username string, password string) (user models.User)
	GetUserByUsername(username string) (params models.Params)
	Save(user models.User) (int, models.User)
	Deletes(username string) (user models.User)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct{}

//登录
func (n userRepository) GetUserByUserNameAndPwd(username string, password string) (user models.User) {
	db := datasource.GetDB()
	db.Where("username = ? and password = ?", username, password).First(&user)
	return
}

func (n userRepository) GetUserByUsername(username string) (param models.Params) {
	db := datasource.GetDB()
	rows, err := db.Table("user").Select("user.*, type.name as type_name, type.code as type_code").Joins("left join type on type.id = user.type_id").Rows()

	if err != nil {
		return nil
	}

	params, errs := models.FormatResult(rows)

	if errs != nil {
		return nil
	}

	param = params[0]
	delete(param, "password")

	return
}

func (n userRepository) Deletes(username string) models.User {
	var user models.User
	user.Username = username
	db := datasource.GetDB()
	db.Where("username = ?", username).Delete(&user)
	return user
}

//添加/修改
func (n userRepository) Save(user models.User) (int, models.User) {
	code := 0
	db := datasource.GetDB()
	if user.ID != 0 {
		var oldUser models.User
		datasource.GetDB().First(&oldUser, user.ID)
		user.CreatedAt = oldUser.CreatedAt
		user.Username = oldUser.Username
		if user.Name == "" {
			user.Name = oldUser.Name
		}
		if user.Email == "" {
			user.Email = oldUser.Email
		}
		if user.Mobile == "" {
			user.Mobile = oldUser.Mobile
		}
		if user.QQ == "" {
			user.QQ = oldUser.QQ
		}
		if user.Gender == 0 {
			user.Gender = oldUser.Gender
		}
		if user.Age == 0 {
			user.Age = oldUser.Age
		}
		if user.Remark == "" {
			user.Remark = oldUser.Remark
		}
		if err := db.Save(&user).Error; err != nil {
			log.Println("报错了==>", err)
			code = -1
		}
	} else {
		if user.Password != "" {
			user.Password = utils.GetMD5String(user.Password)
		}
		if err := db.Create(&user).Error; err != nil {
			log.Println("报错了==>", err)
			code = -1
		}
	}

	return code, user
}
