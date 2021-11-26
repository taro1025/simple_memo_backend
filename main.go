package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"
	"simple_memo/controller"
	"time"

	//"jwt_work/register"
	//"jwt_work/login"
	//"jwt_work/user"
)
//gin
//https://github.com/gin-gonic/gin#gin-v1-stable

func setupRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	//TODO railsみたいにpathを返すヘルパーを使いたい
	v1 := router.Group("/v1")
	{
		memo := v1.Group("/memos")
		{
			memo.GET("/index", controller.Index)
			memo.POST("/create", controller.CreateMemo)
		}
		//auth := v1.Group("/auth")
		{
			//auth.POST("/login", controller.Login)
			//auth.GET("/logout", controller.Logout)
		}
	}
	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}