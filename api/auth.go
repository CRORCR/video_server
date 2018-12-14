package main

import (
	"net/http"
	"video_server/api/defs"
	"video_server/api/session"
)

/**
 * @desc    健全 校验之类  身份验证
 * @author Ipencil
 * @create 2018/12/14
 */
var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

//校验是否是合法的session
func validateUserSession(r *http.Request) bool {
	//不存在,返回
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 { //sid为空,返回false
		return false
	}
	//存在,是否过期
	uname, ok := session.IsSessionExpired(sid)
	if ok { //是否过期
		return false
	}
	//存在,没有过期
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
