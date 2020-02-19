package datasource

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {

	var err error
	var urlStr string
	databasetype := os.Getenv("DATABASETYPE")
	user := os.Getenv("DATABASEUSER")
	password := os.Getenv("DATABASEPASSWORD")
	host := os.Getenv("DATABASEHOST")
	port := os.Getenv("DATABASEPORT")
	tablename := os.Getenv("DATABASETABLENAME")
	urlStr = user + ":" + password + "@tcp(" + host + ":" + port + ")/" + tablename + "?charset=utf8&parseTime=True&loc=Local"

	fmt.Println("mysqp=>", databasetype, "----", urlStr)
	db, err = gorm.Open(databasetype, urlStr)

	if err != nil {
		fmt.Println(err)
		panic("数据库连接失败")
	}

	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(1 * time.Second)
	db.DB().SetMaxIdleConns(20)   //最大打开的连接数
	db.DB().SetMaxOpenConns(2000) //设置最大闲置个数
	db.SingularTable(true)        //表生成结尾不带s
	// 启用Logger，显示详细日志
	db.LogMode(true)
	Createtable()
}
