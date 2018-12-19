package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/**
 * @desc    streaming 上传和下载,肯定是长连接,和http的短连接不同
 * @author Ipencil
 * @create 2018/12/14
 */

//多个长连接同时访问,肯定要有流控

//注册route
func RegisterHanlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos:/vid-id", streamHandler)
	router.POST("/upload:/vid-id", uploadHandler)
	return router
}

func main() {
	r := RegisterHanlers()
	mh := NewMiddleWareHandler(r,2) //中间件 注册到server中  流控制随意设置
	//http.ListenAndServe(":9000", r)
	http.ListenAndServe(":9000", mh)
}

//中间处理器
type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}


func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn(){ //能否获取连接,如果能获取就在桶内放一个数据,如果连接数够了,就返回false
		sendErrorReponse(w,http.StatusTooManyRequests,"Too Many Requests.")   //tooMany 这个错误码是http规定的,429 超过流控值的意思
		return
	}
	m.r.ServeHTTP(w, r)
	//连接断开,需要把连接放回去(取出chan中一个元素)
	defer m.l.ReleaseConn()
}
