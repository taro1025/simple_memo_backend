package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestIndexMemo(test *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/memo/index", nil)
	router.ServeHTTP(w, req)
	assert.Equal(test, 200, w.Code)
}

func TestCreateMemo(test *testing.T) {
	router := setupRouter()
	data := url.Values{"text": {"bar"}}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/memo/create", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(test, 201, w.Code)
}
