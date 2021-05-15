package tran

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id   int
	Name string
	Age  int
	Addr string
}

type Getdata struct {
	Name string
	Age  int
}

type UserProfile struct {
	Id     int
	Pic    string
	CPic   string
	Phone  string
	UserID int // 默认关联字段为Id
	Name   string
	Age    int
	Addr   string
}

func tran(context *gin.Context) {

	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	defer db.Close() // 关闭连接

	if nil != err {
		fmt.Println(err)
	}

	ct := db.Begin() // 开启事务

	var users []User
	db.Where("age = ?", 100).Find(&users).Update("name", "hi")

	ret3 := ct.Commit()

	if ret3.Error != nil {
		fmt.Println("aa")
		ct.Rollback() // 回滚
	}
}
