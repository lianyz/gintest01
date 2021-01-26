package main

import (
	"encoding/json"
	"fmt"
	"github.com/liuhongdi/gintest01/router"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"github.com/liuhongdi/gintest01/pkg/unittest"
	"net/http/httptest"
	"testing"
)

type ResultCont struct {
	Code	int  `json:"code"` // 自增
	Msg string  `json:"msg"` //
	Data interface{} `json:"data"`
}

//测试得到用户信息
func TestGet(t *testing.T) {
	globalrouter = router.Router()
	var w *httptest.ResponseRecorder

	urlIndex := "/user/get?uid=3"
	w = unittest.Get(urlIndex, globalrouter)

	assert.Equal(t,200, w.Code)
	//result := w.Body
	body,_ := ioutil.ReadAll(w.Body)

	m:=&ResultCont{}

	err := json.Unmarshal(body, &m)
	if err != nil {

		fmt.Println("Umarshal failed:", err)
		return
	}
	assert.Equal(t,0, m.Code)
}
//测试post登录
func TestLogin(t *testing.T) {
	globalrouter = router.Router()
	var w *httptest.ResponseRecorder
	//test post
	param := make(map[string]string)
	param["username"] = "username"
	param["password"] = "password"
	urlLogin := "/user/post"
	w = unittest.PostForm(urlLogin, param, globalrouter)
	assert.Equal(t,200, w.Code)
	body,_ := ioutil.ReadAll(w.Body)
	m:=&ResultCont{}

	err := json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}

	if (m.Code != 0){
		t.Fatalf("fail,code:%d,msg:%s",m.Code,m.Msg)
	}
}

