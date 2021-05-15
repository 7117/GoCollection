package logs_source

import (
	"encoding/json"
	"gin_project/util"
	"io/ioutil"
	"os"
)

type LogConfig struct {
	LogDir   string `json:"log_dir"`
	LogLevel string `json:"log_level"`
	LogName  string `json:"log_name"`
}

func LoadLogConfig() *LogConfig {
	log_conf := LogConfig{}

	file, err := os.Open("config/log.json")
	util.Err(err)
	defer file.Close()

	byte_data, err := ioutil.ReadAll(file)
	util.Err(err)

	err = json.Unmarshal(byte_data, &log_conf)


	Log.Info(&log_conf)
	Log.Info(log_conf)
	util.Err(err)

	return &log_conf

}
