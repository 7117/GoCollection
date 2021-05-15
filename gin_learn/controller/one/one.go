package one

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

	db.AutoMigrate(&UserProfile{})

	//插入
	//user_profile := UserProfile{
	//	Pic:   "1.jpg",
	//	CPic:  "2.jpg",
	//	Phone: "xxx",
	//	User:  User{Name: "hallen", Age: 18, Addr: "xxx"},
	//}
	//
	//result := db.Create(&user_profile)

	// 查询1
	// 先查询出来，在根据Association属性值关联查询，Association的值为关联表的模型名称
	//var u_profile UserProfile
	//db.First(&u_profile, 1)
	//fmt.Println(u_profile)
	//db.Model(&u_profile).Association("User").Find(&u_profile.User)
	//fmt.Println(u_profile)

	//查询2
	//var u_profile UserProfile
	//db.Preload("User").First(&u_profile, 1) // 关系名
	//fmt.Println(u_profile)

	//查询3
	//var user_profile3 UserProfile
	//db.First(&user_profile3, 1)
	//var user User
	//db.Model(&user_profile3).Related(&user, "User")
	//fmt.Println(user)
	//fmt.Println(user_profile3)

	//更新
	//var u_profile1 UserProfile
	//db.Debug().Preload("User").First(&u_profile1,1)
	//fmt.Println(u_profile1)
	//db.Model(&u_profile1.User).Update("name","aaa")

	//删除 通过主表删除关联表中的
	//var u_profile2 UserProfile
	//db.Debug().Preload("User").First(&u_profile2, 1)
	//db.Delete(&u_profile2.User)

}
