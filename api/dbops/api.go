package dbops

import (
	_ "github.com/go-sql-driver/mysql"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2018/12/9
 */
func AddUserCredential(userName, pwd string) error {
	//Prepare 预编译
	stmtIns,err:=dbConn.Prepare("insert into user (user_name,pwd) values (?,?)")
	defer func() {stmtIns.Close()}()
	if err!=nil{
		return err
	}
	_, err = stmtIns.Exec(userName, pwd)
	if err!=nil{
		return err
	}

	return nil
}

func GetUserCredenttail(userName string) string {
	stmtOut,er:=dbConn.Prepare("select pwd from user where user_name=?")
	defer func() {stmtOut.Close()}()
	if er!=nil{
		return ""
	}
	var pwd string
	stmtOut.QueryRow(userName).Scan(&pwd)
	return pwd
}

func DeleteUser(userName,pwd string)error{
	stmtDel,err:=dbConn.Prepare("delete from user where user_name=? and pwd=?")
	defer func() {stmtDel.Close()}()
	if err!=nil{
		return err
	}
	_, err = stmtDel.Exec(userName, pwd)
	if err!=nil{
		return err
	}
	return nil
}