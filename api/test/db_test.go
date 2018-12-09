package test

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gpmgo/gopm/modules/log"
)

const (
	dataSourceName = "root:root@tcp(127.0.0.1:3306)/test?charset=utf8"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
}

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2018/12/8
 */

type User struct {
	Name     string `json:"name" xorm:"name"`
	Password string `json:"password" xorm:"password"`
}

/**
 * @desc   : insert
 * @author : Ipencil
 * @date   : 2018/12/8
 */
func (u *User) Insert() error {
	_, e := engine.Table("user").Insert(u)
	return e
}

func TestVV(t *testing.T) {
	u:=&User{Name:"lcq"}
	_, e := engine.Table("user").Get(u)
	fmt.Println("get user", e)
	fmt.Println(u)
}
