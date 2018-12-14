package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//中间件处理器    实现serverHTTP
type middleWareHandler struct {
	r *httprouter.Router
}

//实现http.Handle接口   检查session后调用serverHttp
//包一层,请求来了之后,先经过这个处理器
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//检查session合法性
	validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := &middleWareHandler{}
	m.r = r
	return m
}

//httprouter.Router 实际上实现了 go里面的http.Handle接口
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:name", Login)
	//router.POST("/register:name/pass:pass", Register)
	return router
}

/**
 * @desc    用户操作
 * @author Ipencil
 * @create 2018/12/7
 */
func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
