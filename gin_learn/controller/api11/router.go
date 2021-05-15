package api11

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.GET("/axios", Create)
	router.GET("/detail", Detail)
}
