package data_source

import (
	"encoding/json"
	"gin_project/util"
	"io/ioutil"
	"os"
)

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	LogoMode bool   `json:"logo_mode"`
}

func LoadMysqlConf() *MysqlConf {

	mysql_conf := MysqlConf{}

	file, err := os.Open("config/mysql.json")

	util.Err(err)

	defer file.Close()

	byte_data, err3 := ioutil.ReadAll(file)
	util.Err(err3)

	err4 := json.Unmarshal(byte_data, &mysql_conf)
	util.Err(err4)

	return &mysql_conf

}
