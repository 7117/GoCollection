package user

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func User(context *gin.Context) {
	context.String(http.StatusOK, "User")
}
