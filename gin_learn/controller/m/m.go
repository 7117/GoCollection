package m

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User2 struct {
	Id   int
	Name string
	Age  int
	Addr string
	//不论多对一还是一对多  第一个参数都是子表的  第二个参数是主表的
	//多对一：uid是对面的   id是自身的
	Article []Article `gorm:"ForeignKey:UId;AssociationForeignKey:Id"`
}

type Article struct {
	Id      int
	Title   string
	Content string
	Desc    string
	UId     int
	//第一个指示是：自身的
	//第二个指示是：外部的
	//User2 []User2 `gorm:"ForeignKey:UId;AssociationForeignKey:Id"`
}

func M(context *gin.Context) {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close() // 关闭连接

	if nil != err {
		fmt.Println(err)
	}

	//add(db)
	//show(db)
	//updata(db)
	//del(db)

}

func del(db *gorm.DB) {
	// 先查询
	var user2 User2
	//db.Preload("Articles").Find(&user2, 2) // 关系名
	//// 再删除，删除要指定条件，不然会把所有满足条件的都删除
	//db.Delete(&user2.Article, "title=? and uid=?", "标题测试3", 2)

	// 或者使用where
	db.Where("title=? and uid=?", "标题测试3", 2).Delete(&user2.Article)
}

func updata(db *gorm.DB) {
	// 先查询
	var user2 User2

	db.Preload("Article").Find(&user2, 2) // 关系名

	// 再更新，更新指定条件，不然会把所有满足条件的都更新
	db.Model(&user2.Article).Where("title=? and uid=?", "标题测试2", 2).Update("title", "标题测试3") // name和uid限制条件
}

func show(db *gorm.DB) {
	var user2 User2

	db.Preload("Article").Find(&user2, 2) // 关系名
	fmt.Println(user2)
}

func add(db *gorm.DB) {
	user := User2{
		Name: "halle",
		Age:  18,
		Addr: "xxx",
		Article: []Article{
			{
				Title:   "标题测试",
				Content: "内容测试",
				Desc:    "描述测试",
			}, {
				Title:   "标题测试2",
				Content: "内容测试2",
				Desc:    "描述测试2",
			},
		},
	}

	ret := db.Create(&user)

	fmt.Println(ret.RowsAffected)
	fmt.Println(ret.Error)
}
