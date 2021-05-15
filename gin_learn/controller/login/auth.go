package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var map_data map[string]interface{} = map[string]interface{}{
	"zs": gin.H{"age": 18, "addr": "addrss"},
	"ws": gin.H{"age": 18, "addr": "addrww"},
}

func Auth(context *gin.Context) {
	user_name := context.Query("user_name")
	data, ok := map_data[user_name]

	if ok {
		context.JSON(http.StatusOK, gin.H{
			"user_name": user_name,
			"data":      data,
		})
	} else {
		context.JSON(http.StatusNotFound, gin.H{
			"user_name": user_name,
			"data":      data,
		})
	}

}
