package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"my_gin_cli/conf"
	"my_gin_cli/controller"
	"my_gin_cli/logger"
	"my_gin_cli/router"
	"my_gin_cli/serializer"
	"net/http"
	"net/http/httptest"
	"testing"
)

var r *gin.Engine
func init()  {
	logger.LogConf()
	conf.Init()
	r = router.NewRouter()
}

var cookie *http.Cookie

func TestUserLogin(t *testing.T)  {
	user := &controller.UserLoginService{
		UserName: "admin",
		Password: "123456",
	}

	url := "/api/v1/user/login"
	mashal, _ := json.Marshal(user)
	// 通过NewRecorder()函数构造响应
	res := httptest.NewRecorder()
	// 通过NewRequest构造请求
	req := httptest.NewRequest("POST", url, bytes.NewBuffer(mashal))

	req.Header.Add("content-type", "application/json")

	r.ServeHTTP(res, req)

	// 提取响应
	result := res.Result()
	defer result.Body.Close()
	cookie = result.Cookies()[0]
	// 读取响应body
	body,_ := ioutil.ReadAll(result.Body)
	fmt.Printf(string(body))
	json.Unmarshal(body, &serializer.Response{})
}

func TestUserList(t *testing.T)  {
	url := "/api/v1/user/list"
	req := httptest.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()


	req.Header.Add("content-type", "application/json")
	req.AddCookie(cookie)

	r.ServeHTTP(resp, req)

	result := resp.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &serializer.Response{})
}



func TestMyFoo(t *testing.T)  {
	t.Run("user=login", TestUserLogin)
	t.Run("user=list", TestUserList)
	t.Run("user=register", TestUserRegister)
}