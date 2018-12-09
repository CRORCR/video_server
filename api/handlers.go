package main

import (
	"fmt"
	"io"
	"net/http"
	"video_server/api/server"

	"github.com/julienschmidt/httprouter"
)

 func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
 	io.WriteString(w,"hello world")
 	return
 }

 func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	 name := p.ByName("name")
	 fmt.Println("name",name)
	 user:=&server.User{Name:name}
	 err := user.GetUser()
	 if err!=nil{
		 io.WriteString(w,err.Error())
	 }
	 io.WriteString(w,"success")
	 return
 }

func Register(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	name := p.ByName("name")
	pass := p.ByName("pass")
	fmt.Println("name",name)
	user:=&server.User{name,pass}
	err := user.Insert()
	if err!=nil{
		io.WriteString(w,err.Error())
	}
	io.WriteString(w,"success")
	return
}