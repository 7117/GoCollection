package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {

	router := httprouter.New()

	// 主页
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)

	// 主页
	router.GET("/home", homeHandler)
	router.POST("/home", homeHandler)

	// 用户主页
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)

	// api代理
	router.POST("/api", apiHandler)

	// proxy
	router.POST("/upload/:vid-id", proxyHandler)

	// 静态
	// 会自动挂载  把tempalte文件夹下面的东西挂载到statics
	// 127.0.0.1:8080/statics/(template文件夹下面的内容)
	router.ServeFiles("/statics/*filepath",http.Dir("./template"))

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}
