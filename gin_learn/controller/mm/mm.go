package mm

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Article2 struct {
	AId     int `gorm:"primary_key:true"`
	Title   string
	Content string
	Desc    string
	Tags    []Tag `gorm:"many2many:Article2s2Tags" ;ForeignKey:AId;AssociationForeignKey:TId`
	//Tags    []Tag `gorm:"many2many:Article2s2Tags;ForeignKey:AId;AssociationForeignKey:TId"`
}

type Tag struct {
	TId  int `gorm:"primary_key:true"`
	Name string
	Desc string
}

func Mm(context *gin.Context) {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if nil != err {
		fmt.Println(err)
	}

	defer db.Close() // 关闭连接

	//add(db)
	show(db)
	//updata(db)
	//del(db)
}

func del(db *gorm.DB) {
	// 先查询
	var article Article2
	// 再删除，记得加条件
	db.Where("name = ?", "xxx").Delete(&article.Tags)
}

func updata(db *gorm.DB) {
	// 先查询
	var article Article2
	db.Preload("Tags").Find(&article, 1)
	// 再更新，记得加条件
	db.Model(&article.Tags).Where("t_id = ?", 2).Update("name", "xxx")
}

func show(db *gorm.DB) {
	var article Article2

	db.Preload("Tags").Find(&article, 1)
	fmt.Println(article)
}

func add(db *gorm.DB) {
	// 一篇文章有多个帖子
	article := Article2{
		Title:   "测试多对多标题1",
		Content: "测试多对多内容1",
		Desc:    "测试多对多描述1",
		Tags: []Tag{
			{
				Name: "django",
				Desc: "django标签",
			},
			{
				Name: "python",
				Desc: "python标签",
			},
		},
	}

	ret := db.Create(&article)
	fmt.Println(ret.Error)
}
