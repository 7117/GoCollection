package log

import (
	"gin_project/logs_source"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func LogRecord(context *gin.Context) {
	logs_source.Log.Info("info  info")

	logs_source.Log.WithFields(logrus.Fields{
		"msg": "测试的错误",
		"aaa": []int{1, 2, 3, 4, 5},
	}).Warn("这是一个warnning级别的日志")
}
