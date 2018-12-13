package defs

/**
 * @desc    数据模型
 * @author Ipencil
 * @create 2018/12/9
 */
//request
 type UserCredential struct{
 	Name string `json:"name"`
 	Pwd string `json:"pwd"`
 }

//video info
type VideoInfo struct{
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Username string //login name
	TTL int64
}
