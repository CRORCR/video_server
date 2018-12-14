package main

import (
	"io"
	"net/http"
)

/**
 * @desc    返回数据
 * @author Ipencil
 * @create 2018/12/14
 */

 //只有出错了,才返回
 func sendErrorReponse(w http.ResponseWriter,sc int,errMsg string){
	w.WriteHeader(sc)
	io.WriteString(w,errMsg)
 }