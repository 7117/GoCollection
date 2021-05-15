package err

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

func Hello(context *gin.Context) {

	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	defer db.Close() // 关闭连接

	if nil != err {
		fmt.Println(err)
	}

	GetError(db)
	//GetErrors(db)

}

func GetError(db *gorm.DB) {
	var user User

	result := db.Where("name = ?", "jinzhasu").First(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println("aaaa")
}

func GetErrors(db *gorm.DB) {
	var user User

	errors := db.Where("name = ?", "jinzhasu").First(&user).GetErrors()
	fmt.Println(len(errors)) // 打印错误数量

	// 遍历错误内容
	for _, err := range errors {
		fmt.Println(err)
	}

	fmt.Println("bbbb")

}
