package raw

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	{
		router.GET("/raw", One)
	}
}
