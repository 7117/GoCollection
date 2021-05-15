package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MiddlewareA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MiddlewareA before request")
		// before request
		c.Next()
		// after request
		fmt.Println("MiddlewareA after request")
	}
}
