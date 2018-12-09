package main

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2018/12/9
 */
import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter){
	io.WriteString(w,"")
}


func sendNormalResponse(w http.ResponseWriter){

}