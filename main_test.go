package main

import (
	"bytes"
	jwt "github.com/appleboy/gin-jwt/v2"
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
	loginBefore()
	code := m.Run()
	//after()
	os.Exit(code)
}

var validLoginInfo, invalidLoginInfo *bytes.Buffer
var authMiddleware *jwt.GinJWTMiddleware


func loginBefore() {
	validLoginInfo = bytes.NewBufferString("{\"email\":\"rintaro.sakino@msetsu.com\", \"password\":\"password\"}")
	invalidLoginInfo = bytes.NewBufferString("{\"email\":\"rintetsu.com\", \"password\":\"\"}")
	authMiddleware = middleware.Auth()
}

type loginType struct {
	title  string
	body   *bytes.Buffer
	status int
}

//TODO テストデータ用意したい
func TestLogin(test *testing.T) {
	var loginTests = []loginType{
		{"異常:Passwordがないので", invalidLoginInfo, 401},
		{"正常:正しい", validLoginInfo, 200},
	}
	for _, tt := range loginTests {
		test.Run(tt.title, func(t *testing.T) {
			_ = login(test, tt.body, tt.status)
		})
	}
}

func login(test *testing.T, body *bytes.Buffer, expectStatus int) *gin.Context {
	router := setupRouter()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/login", body)
	c.Request.Header.Add("Content-Type", binding.MIMEJSON)
	authMiddleware.LoginHandler(c)
	router.ServeHTTP(w, c.Request)
	if w.Result().StatusCode != expectStatus {
		test.Errorf(`Login: expect "%v", but "%v"`, expectStatus, w.Code)
	}
	return c
}
