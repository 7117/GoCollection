package logs_source

import (
	"gin_project/util"
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	// 日志級別映射
	log_level_mapping := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fata":  logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}

	log_conf := LoadLogConfig()

	dir, err := os.OpenFile(log_conf.LogDir+"/"+log_conf.LogName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	util.Err(err)

	Log.Out = dir
	Log.Level = log_level_mapping[log_conf.LogLevel]

	//设置日志格式，格式化时间
	Log.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
}
