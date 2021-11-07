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
	if err != nil{
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	memoService :=service.MemoService{}
	err = memoService.SetMemo(&memo)
	if err != nil{
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func Index(c *gin.Context) {
	memoService :=service.MemoService{}
	BookLists := memoService.Index()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data": BookLists,
	})
}
