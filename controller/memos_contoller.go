package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_memo/model"
	"simple_memo/service"
)

func CreateMemo(c *gin.Context) {
	memo := model.Memo{}
	err := c.Bind(&memo)
	user, _ := c.Get("id")
	if err != nil{
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	memoService :=service.MemoService{Db: service.Db}
	validUser, ok := user.(*model.User)
	if !ok {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	err = memoService.SetMemo(validUser, &memo)
	if err != nil{
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func Index(c *gin.Context) {
	memoService :=service.MemoService{Db: service.Db}
	memos := memoService.Index()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data": memos,
	})
}
