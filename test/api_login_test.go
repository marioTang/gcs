package test

import (
	"encoding/json"
	"gcs/utils/resp"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"testing"
)

var (
	Token    = g.MapStrStr{}
	Username = "admin"
)

func TestSystemConfigIndex(t *testing.T) {
	// 登录，访问用户信息
	if r, e := ghttp.Get(TestURL + "/system/config/index"); e != nil {
		t.Error(e)
	} else {
		content := string(r.ReadAll())
		if r.StatusCode != 200 {
			t.Error("content error", content)
		} else if gstr.Contains(content, "code") && gstr.Contains(content, "msg") {
			t.Error("content code error", content)
		} else {
			t.Log("index page no auth")
		}

		r.Close()
	}

}

func TestSystemUserGet(t *testing.T) {
	// 登录，访问用户信息
	t.Log("1. execute get user")
	data := Post(t, "/system/user/get", "id=1")
	if data.Success() {
		t.Log(data.Json())
	} else {
		t.Error(data.Json())
	}
}

func TestLogin(t *testing.T) {
	Username = "testLogin"
	t.Log(" login first ")
	token1 := getToken(t)
	t.Log("token:" + token1)
	t.Log(" login second and same token ")
	token2 := getToken(t)
	t.Log("token:" + token2)
	if token1 != token2 {
		t.Error("token not same ")
	}
	Username = "flyfox"
}

func TestLogout(t *testing.T) {
	Username = "testLogout"
	t.Log(" logout test ")
	data := Post(t, "/user/logout", "username="+Username)
	if data.Success() {
		t.Log(data.Json())
	} else {
		t.Error(data.Json())
	}
	Username = "flyfox"
}

func Post(t *testing.T, urlPath string, data ...interface{}) resp.Resp {
	client := ghttp.NewClient()
	client.SetHeader("Authorization", "Bearer "+getToken(t))
	content := client.RequestContent("POST", TestURL+urlPath, data...)
	var respData resp.Resp
	err := json.Unmarshal([]byte(content), &respData)
	if err != nil {
		t.Error(err)
	}
	return respData
}

func getToken(t *testing.T) string {
	if Token[Username] != "" {
		return Token[Username]
	}

	passwd, _ := gmd5.Encrypt("123456")
	if r, e := ghttp.Post(TestURL+"/login/submit", "username="+Username+"&passwd="+passwd); e != nil {
		t.Error(e)
	} else {
		defer r.Close()

		content := string(r.ReadAll())
		var respData resp.Resp
		err := json.Unmarshal([]byte(content), &respData)
		if err != nil {
			t.Error(err)
		}

		if !respData.Success() {
			t.Error("resp fail:" + respData.Json())
		}

		Token[Username] = respData.GetString("token")
	}
	return Token[Username]
}
