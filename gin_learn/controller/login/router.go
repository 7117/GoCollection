package login

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	{
		router.GET("/auth", gin.BasicAuth(gin.Accounts{
			"zs": "000",
			"ls": "123",
			"ww": "456",
		}), Auth)
	}
}
