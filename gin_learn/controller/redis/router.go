package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("hallen"))

	router.Use(sessions.Sessions("session_test", store))

	{
		router.GET("/redis", RedisTest)
	}
}
