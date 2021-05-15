package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionTest(context *gin.Context) {
	// 初始化session对象
	session := sessions.Default(context)

	// 设置session
	session.Set("name", "hallen")

	// 获取session
	//name := session.Get("name")
	//fmt.Println(name)

	//// 删除指定key的session
	//session.Delete("name")
	//
	//name2 := session.Get("name")
	//fmt.Println("================++++++++++")
	//fmt.Println(name2)
	//
	//fmt.Println(session.Get("age"))
	//// 删除所有的session
	//session.Clear()
	//fmt.Println(session.Get("age"))
	//fmt.Println(session.Get("addr"))
	//
	//// 保存session
	//session.Save()
}
