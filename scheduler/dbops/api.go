package dbops

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
)

/**
 * @desc    任务接收,写入数据库
 * @author Ipencil
 * @create 2018/12/22
 */
func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES(?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}