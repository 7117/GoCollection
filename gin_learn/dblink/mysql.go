package dblink

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB
var err error

type User struct {
	Id   int
	Name string
	Age  int
	Addr string
}

func init() {
	var user User

	Db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if nil != err {
		fmt.Println(err)
		panic(err)
	}

	//Db.DB().SetMaxOpenConns(100) // 最大连接数
	//Db.DB().SetMaxIdleConns(50) // 最大空闲数
	defer Db.Close() // 关闭连接

	Db.AutoMigrate(&user)
}
