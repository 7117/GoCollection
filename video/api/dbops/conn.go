package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

)

// 定义全局变量
var (
	dbConn *sql.DB
	err error
)


func init() {
	dbConn,err = sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/video")

	if err != nil {
		panic(err.Error())
	}

}