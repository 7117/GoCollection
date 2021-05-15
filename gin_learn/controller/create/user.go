package create

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string
	Age      int
	AnimalId int64 `gorm:"column:beast_id"`
}

func Create(context *gin.Context) {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if nil != err {
		fmt.Println(err)
	}

	defer db.Close() // 关闭连接

	db.CreateTable(&User{}) // 不指定表名,模型后面会加个s

}
