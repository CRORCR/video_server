package streamserver

import "github.com/julienschmidt/httprouter"

/**
 * @desc    streaming 上传和下载,肯定是长连接,和http的短连接不同
 * @author Ipencil
 * @create 2018/12/14
 */

//多个长连接同时访问,肯定要有流控

//注册route
func RegisterHanlers() *httprouter.Router {

}

func main() {

}