package dbops

import (
	"testing"
)

//一般test文件有一个init
//这里同文件夹有了初始化init数据库了,这里就不用提供
/**
 * @desc    数据库相关测试
 * @author Ipencil
 * @create 2018/12/9
 */
func clearTable() {
	dbConn.Exec("truncate user")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate sessions")
	dbConn.Exec("truncate comments")
}

func TestMain(m *testing.M) {
	clearTable()
	m.Run()
	clearTable()
}


func TestUserWorkFlow(t *testing.T) {
	t.Run("add",testAddUser)   //添加一个用户
	t.Run("get",testGetUser)   //查询这个用户
	t.Run("del",testDelteUser) //删除这个用户
	t.Run("reg",testRegetUser) //再次查询这个用户
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("lcq", "123")
	if err != nil {
		t.Errorf("Error of add user%v\n", err)
	}
}

func testGetUser(t *testing.T) {
	pwd := GetUserCredenttail("lcq")
	if pwd == "" {
		t.Errorf("Error of get user%v\n", err)
	}
}

func testDelteUser(t *testing.T) {
	err := DeleteUser("lcq", "123")
	if err != nil {
		t.Errorf("Error of delete user%v\n", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd := GetUserCredenttail("lcq")
	if pwd != "" {
		t.Errorf("Error of get user%v\n", pwd)
	}
}