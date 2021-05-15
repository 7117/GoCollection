package middle

import "github.com/gin-gonic/gin"

func GetCrosQues(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")

	context.Next()
}
