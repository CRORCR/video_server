package server

import (
	_ "github.com/go-sql-driver/mysql"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2018/12/8
 */

type User struct {
	UserName string `json:"name" xorm:"user_name"`
	Password string `json:"password" xorm:"password"`
}
