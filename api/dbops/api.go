package dbops

import (
	"database/sql"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"

	_ "github.com/go-sql-driver/mysql"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2018/12/9
 */
func AddUserCredential(userName, pwd string) error {
	//Prepare 预编译
	stmtIns, err := dbConn.Prepare("insert into user (user_name,pwd) values (?,?)")
	defer func() { stmtIns.Close() }()
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(userName, pwd)
	if err != nil {
		return err
	}

	return nil
}

func GetUserCredenttail(userName string) string {
	stmtOut, er := dbConn.Prepare("select pwd from user where user_name=?")
	defer func() { stmtOut.Close() }()
	if er != nil {
		return ""
	}
	var pwd string
	stmtOut.QueryRow(userName).Scan(&pwd)
	return pwd
}

func DeleteUser(userName, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from user where user_name=? and pwd=?")
	defer func() { stmtDel.Close() }()
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(userName, pwd)
	if err != nil {
		return err
	}
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//创建uuid 在utils包内
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	ctime := time.Now().Format("Jan 02 2006,15:04:05")
	stmtIns, err := dbConn.Prepare(`insert into video_info 
		(id, author_id, name, display_ctime) values(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	defer func() { stmtIns.Close() }()
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	return res, nil
}

//获得一个视频,根据视频id获得
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	defer func() { stmtOut.Close() }()
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	defer func() { stmtDel.Close() }()
	if err != nil {
		return err
	}
	return nil
}

//添加一个评论 评论的信息 只能添加和查询   不能修改和删除
//需要用户id 视频id 内容content 三个内容
func AddNewComments(vid string, aid int, content string) error {
	//获得uuid
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

//	  left join(左联接)   返回包括左表中的所有记录和右表中联结字段相等的记录
//　　 right join(右联接)  返回包括右表中的所有记录和左表中联结字段相等的记录
//　　inner join(等值连接) 只返回两个表中联结字段相等的行
//获得所有的评论   user表和comments表做关联  使用内连接

//FROM_UNIXTIME  只会精确到秒,所以用左开右闭区间,这样可以得到最后(当前)插入的数据
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, user.user_name, comments.content FROM comments
		INNER JOIN user ON comments.author_id = user.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}
