package models

type User struct {
	Username string `gorm:"unique;unique_index" json:"username"`
	Password string `json:"password"`
	TypeId   int    `json:"type_id"`
	Name     string `json:"name"`   //姓名
	Email    string `json:"email"`  //邮箱
	Mobile   string `json:"mobile"` //手机
	QQ       string `json:"qq"`
	Gender   int    `json:"gender"` //0男 1女
	Age      int    `json:"age"`    //年龄
	Remark   string `json:"remark"` //备注
	Token    string `gorm:"-"`
	Session  string `gorm:"-"`
	Model
}
