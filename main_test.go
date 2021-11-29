package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"simple_memo/controller"
	"testing"
)

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	//after()
	os.Exit(code)
}

func before() {
	//service.DbEngine
}

func TestLogin(test *testing.T) {
	//user := model.User{Email: "test@test.com", Password: "password"}
	//userForm, _ := json.Marshal(user)
	body := bytes.NewBufferString("{\"email\":\"se@tst.om\", \"password\":\"password\"}")
	fmt.Println(body)
	//router := setupRouter()
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	ginContext.Request, _ = http.NewRequest("POST", "/v1/users/create", body)
	ginContext.Request.Header.Add("Content-Type", binding.MIMEJSON)

	controller.CreateUser(ginContext)
	//router.ServeHTTP(w, c.Request)
	assert.Equal(test, 201, w.Code)
}
//
//func TestCreateUser(t *testing.T) {
//	userTests := []struct {
//		user model.User
//	}{
//		{model.User{Email: "test.com", Password: "password"}},
//		{model.User{Email: "momo@com", Password: "njnjn"}},
//	}
//
//	for _, test := range userTests {
//		w := httptest.NewRecorder()
//		ginContext, _ := gin.CreateTestContext(w)
//		userForm, _ := json.Marshal(test.user)
//		body := bytes.NewBuffer(userForm)
//		ginContext.Request, _ = http.NewRequest("POST", "/v1/users/create", body)
//		ginContext.Request.Header.Add("Content-Type", binding.MIMEJSON)
//		controller.CreateUser(test.user)
//
//	}
//}

//func TestIndexMemo(test *testing.T) {
//	router := setupRouter()
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/v1/memo/index", nil)
//	router.ServeHTTP(w, req)
//	assert.Equal(test, 200, w.Code)
//}
//
//func TestCreateMemo(test *testing.T) {
//	router := setupRouter()
//	//User
//	data := url.Values{"email": {"test@test.com"}, "password": {"password"}}
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/v1/users/create", strings.NewReader(data.Encode()))
//	router.ServeHTTP(w, req)
//	assert.Equal(test, 200, w.Code)
//
//	//Login
//	w = httptest.NewRecorder()
//	req, _ = http.NewRequest("POST", "/v1/login", strings.NewReader(data.Encode()))
//	router.ServeHTTP(w, req)
//	assert.Equal(test, 200, w.Code)
//
//	//Memo
//	data = url.Values{"text": {"bar"}}
//	w = httptest.NewRecorder()
//	req, _ = http.NewRequest("POST", "/v1/memos/create", strings.NewReader(data.Encode()))
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	router.ServeHTTP(w, req)
//	assert.Equal(test, 201, w.Code)
//}

//func TestLogin(test *testing.T) {
//	router := setupRouter()
//	data := url.Values{"email": {"test@test.com"}, "password": {"password"}}
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/v1/users/create", strings.NewReader(data.Encode()))
//	//req, _ := http.NewRequest("POST", "/v1/login", nil)
//	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	router.ServeHTTP(w, req)
//	assert.Equal(test, 200, w.Code)
//
//	w = httptest.NewRecorder()
//	req, _ = http.NewRequest("POST", "/v1/login", strings.NewReader(data.Encode()))
//	//req, _ := http.NewRequest("POST", "/v1/login", nil)
//	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	router.ServeHTTP(w, req)
//	assert.Equal(test, 200, w.Code)
//}
