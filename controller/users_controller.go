package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_memo/model"
	"simple_memo/service"
)

func CreateUser(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil{
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	userService :=service.UserService{Db: service.Db}
	if err = userService.SetUser(&user); err != nil{
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
