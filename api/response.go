package main

/**
 * @desc    处理所有返回消息
 * @author Ipencil
 * @create 2018/12/9
 */
import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

//错误回报
func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC) //错误码写入头信息

	resStr, _ := json.Marshal(&errResp.Error) //错误信息json格式返回
	io.WriteString(w, string(resStr))
}

//成功返回
func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
