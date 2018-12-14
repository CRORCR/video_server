package dbops

import (
	"strconv"
	"testing"
	"time"
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
	//clearTable()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("add", testAddUser)   //添加一个用户
	t.Run("get", testGetUser)   //查询这个用户
	t.Run("del", testDelteUser) //删除这个用户
	t.Run("reg", testRegetUser) //再次查询这个用户
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

var tempvid string

func TestVideoWorkFlow(t *testing.T) {
	clearTable()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

/**
 * @desc   : 评论测试
 * @author : Ipencil
 * @date   : 2018/12/13
 */
func TestComments(t *testing.T) {
	clearTable()
	t.Run("AddUser", testAddUser)           //添加用户
	t.Run("AddCommnets", testAddComments)   //添加评论
	t.Run("ListComments", testListComments) //查询评论
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800 //开始时间
	//int64 转 string 再转int
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10)) //结束时间

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		t.Logf("comment: %d, %v \n", i, ele)
	}
}
