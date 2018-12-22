package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("进入test")
	t, err := template.ParseFiles("K:\\workspace\\src\\video_server\\streamserver\\upload.html")
	if err!=nil{
		fmt.Println("err",err)
		return
	}
	t.Execute(w, nil)
}

/**
 * @desc  上传视频到服务器
 * @author Ipencil
 * @create 2018/12/14
 */
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v", err) //这里错误返回500  内部错误
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4") //设置头信息
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

//视频下载   下载文件大小校验-->从表单from中获取file文件-->读取到本地
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//设置最大读取大小
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	//解析成表单
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}
	//读取表单内容
	file, _, err := r.FormFile("file") // <form name="file" ... >
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}