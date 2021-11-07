package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"simple_memo/controller"
)
//gin
//https://github.com/gin-gonic/gin#gin-v1-stable

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		memo := v1.Group("memo")
		{
			memo.GET("/index", controller.Index)
			memo.POST("create", controller.CreateMemo)
		}
	}
	router.Run(":8080")
}