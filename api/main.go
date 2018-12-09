package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers()*httprouter.Router{
	router:=httprouter.New()
	router.POST("/user",CreateUser)
	router.POST("/login:name",Login)
	router.POST("/register:name/pass:pass",Register)
	return router
}
/**
 * @desc    用户操作
 * @author Ipencil
 * @create 2018/12/7
 */
func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8080",r)
}
