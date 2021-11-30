package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"net/http/httptest"
	"os"
	"simple_memo/middleware"
	"simple_memo/testUtils"
	"testing"
)

func TestMain(m *testing.M) {
	testUtils.Before()
	code := m.Run()
	//after()
	os.Exit(code)
}

type login struct{
	title string
	body *bytes.Buffer
	status int
}

//TODO テストデータ用意したい
func TestLogin(test *testing.T){
	validInfo := bytes.NewBufferString("{\"email\":\"rintaro.sakino@msetsu.com\", \"password\":\"password\"}")
	invalidInfo := bytes.NewBufferString("{\"email\":\"rintetsu.com\", \"password\":\"\"}")
	var loginTests = []login {
		{"異常:Passwordがないので", invalidInfo, 401},
		{"正常:正しい", validInfo, 200},
	}
	for _, tt := range loginTests {
		test.Run(tt.title, func(t *testing.T){
			router := setupRouter()
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest("POST", "/login", tt.body)
			c.Request.Header.Add("Content-Type", binding.MIMEJSON)
			authMiddleware := middleware.Auth()
			authMiddleware.LoginHandler(c)
			router.ServeHTTP(w, c.Request)
			if w.Result().StatusCode != tt.status {
				test.Error("want 401 but", w.Code)
			}
		})
	}
}
