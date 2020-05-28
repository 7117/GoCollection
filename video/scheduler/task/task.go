package task

import (
	"errors"
	"goPractice/video/scheduler/dbops"
	"log"
)

func VideoClearDispatcher(dc dataChan) error {
	res, _ := dbops.ReadVideoDeletionRecord(3)

	if len(res) == 0 {
		return errors.New("剩余0条")
	}

	//写入至数据通道中
	for _, id := range res {
		dc <- id
	}

	return nil
}

//删除实质的文件
func deletVideo(vid string) error {
	err := od.Remove("./videos/" + vid)

	if err != nil {
		log.Print("remove file")
		return err
	}

}

func VideoClearExec(dc dataChan) error {

forloop:
	for {
		select {
		case vid := <-dc:
			go func(vid interface{}) {
				//删除数据库的
				if err := dbops.DeletionRecord(vid.string()); err != nil {
					return
				}
				//删除实质文件的
				if err := deletVideo(vid.(string)); err != nil {
					return
				}
			}(vid)
		default:
			break forloop
		}
	}

}
