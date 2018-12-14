package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/session"

	"github.com/julienschmidt/httprouter"
)

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//使用ioutil原始读取
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	//反序列化成结构体
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	//存储数据库
	if err := dbops.AddUser(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	//存储session 到 map和db
	id := session.GenerateNewSessionId(ubody.Username)
	//返回 success 和 session_id
	su := &defs.SignedUp{Success: true, SessionId: id}
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}
