package dbops

import (
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

)

func AddUserCredential(loginName string,pwd string)error{
	stmtIns,_ := dbConn.Prepare("insert into user (login_name.pwd) values(?,?)")

	stmtIns.Exec(loginName,pwd)
	stmtIns.Close()

	return nil
}

func GetUserCredential(loginName string)(string,error){
	stmtOut,_ := dbConn.Prepare("select pwd from users where login_name = ? ")

	var pwd string

	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()

	return pwd,nil
}

func deleteUser(loginName string,pwd string)error{
	stmtDel,_ := dbConn.Prepare("delete from user where logig_name = ?")

	stmtDel.Exec(loginName)

	stmtDel.Close()

	return nil

}	