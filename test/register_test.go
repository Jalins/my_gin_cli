package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"my_gin_cli/controller"
	"my_gin_cli/serializer"
	"net/http/httptest"
	"testing"
)

func TestUserRegister(t *testing.T)  {
	user := &controller.UserRegisterService{
		Nickname:        "jalins",
		UserName:        "jalins huang",
		Password:        "123456",
		PasswordConfirm: "123456",
	}
	url := "/api/v1/user/register"
	mashalUser, _ := json.Marshal(&user)

	req := httptest.NewRequest("POST", url, bytes.NewBuffer(mashalUser))
	resp := httptest.NewRecorder()

	req.Header.Add("content-type", "application/json")

	r.ServeHTTP(resp, req)

	result := resp.Result()
	defer result.Body.Close()

	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &serializer.Response{})

}