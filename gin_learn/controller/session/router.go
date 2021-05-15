package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	//盐值
	store := cookie.NewStore([]byte("key_salt"))
	router.Use(sessions.Sessions("session_test", store))

	{
		router.GET("/session", SessionTest)
	}
}
