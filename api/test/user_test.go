package test

import (
	"io/ioutil"
	"net/http"
	"testing"
)

/**
 * @desc    用户操作
 * @author Ipencil
 * @create 2018/12/8
 */

var (
	wageArray = []string{
		//"http://localhost:8080/user",
		//"http://localhost:8080/login/lcq",
		"http://localhost:8080/registerlcq/pass1234",
	}
)

func TestCreateUser(t *testing.T) {
	client := &http.Client{}
	for i := 0; i < len(wageArray); i++ {
		query(t, client, wageArray[i])
	}
}

func query(t *testing.T, client *http.Client, url string) {
	reqest, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}
	reqest.Header.Set("user","李长全")
	reqest.Header.Set("password","112233")
	//处理返回
	response, _ := client.Do(reqest)
	defer func() {
		response.Body.Close()}()
	bytes, _ := ioutil.ReadAll(response.Body)
	//解析list
	t.Log("result:", string(bytes))
}
