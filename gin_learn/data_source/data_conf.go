package data_source

import (
	"fmt"
	"gin_project/controller/first"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
var err error

func init() {

	mysql_conf := LoadMysqlConf()

	logo_mode := mysql_conf.LogoMode

	data_source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		mysql_conf.UserName,
		mysql_conf.Password,
		mysql_conf.Host,
		mysql_conf.Port,
		mysql_conf.DataBase,
	)

	Db, err = gorm.Open("mysql", data_source)

	if err != nil {
		panic(err)
	}

	Db.LogMode(logo_mode)

	Db.DB().SetMaxOpenConns(100) // 最大连接数
	Db.DB().SetMaxIdleConns(50)  // 最大空闲数

	Db.AutoMigrate(&first.User{})
}
