package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
/**
 * @desc    数据库连接
 * @author Ipencil
 * @create 2018/12/22
 */
var (
	dbConn         *sql.DB
	err            error
	dataSourceName = "root:root@tcp(127.0.0.1:3306)/video?charset=utf8"
)

//使用init,每次当前包被调用,都会调用init函数
func init() {
	dbConn, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic("DB open failed")
	}
	err = dbConn.Ping()
	if err != nil {
		panic("DB ping failed")
	}
}
