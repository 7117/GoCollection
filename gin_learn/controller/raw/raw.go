package raw

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

type UserProfile struct {
	Id     int
	Pic    string
	CPic   string
	Phone  string
	User   User
	UserID int // 默认关联字段为Id
}

func One(context *gin.Context) {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close() // 关闭连接

	if nil != err {
		fmt.Println(err)
	}

	//selec(db)
	//exe(db)

}

func selec(db *gorm.DB) {
	var users []User
	db.Raw("select * from users").Find(&users)
	fmt.Println(users)
}

func exe(db *gorm.DB) {
	db.Exec("insert into users (name,age) values(?,?)", "hallen222", 111)

}

