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
		"http://localhost:8080/user",
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
	//处理返回
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	//解析list
	t.Log("result:", string(bytes))
}
