package main

import (
	"net/http"
	"video_server/scheduler/taskrunner"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}

func main() {
	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r) //go 需要在listen之前运行,否则go协程都没法启动
}


/*
有两种方式可以阻塞协程
for{} 循环

c:=make(chan int)
...
<-c    定义chan,但是不存入数据,直接取,肯定是阻塞的
 */