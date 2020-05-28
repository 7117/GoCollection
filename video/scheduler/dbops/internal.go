package dbops

import (
	_ "github.com/go-sql-driver/mysql"
)

//查询数据表的
func ReadVideoDeletionRecord(count int)([]string,error){
	stmtOut,_ := dbConn.Prepare("select video_id from video_del_rec limit ?")

	var ids []string

	rows,_ := stmtOut.Query(count)

	for rows.Next(){
		var id string

		if err := rows.Scan(&id);err != nil{
			return ids,err
		}

		ids  = append(ids,id)
	}

	defer stmtOut.close()

	return ids,nil
}

//删除数据表中的
func DeletionRecord(vid string)error{
	stmtDel, _ := dbConn.Prepare("delete from video_del_rec where video_id= ? ")
	defer stmtDel.Close()

	_,err = stmtDel.Exec(vid)

	if err != nil {
		return err
	}

	return nil


}