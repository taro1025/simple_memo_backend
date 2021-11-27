package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"simple_memo/controller"
	"simple_memo/middleware"
	"time"
	//
	//"jwt_work/register"
	//"jwt_work/login"
	//"jwt_work/user"
)

var identityKey = "id"

func main() {
	router := setupRouter()
	router.Run(":8080")
}

func helloHandler(c *gin.Context) {
	if _, err := c.Cookie("jwt"); err != nil {
		log.Printf("no cookie")
	} else {log.Printf("Ok cookie")}
	//claims := jwt.ExtractClaims(c)
	//user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		//"userID":   claims[identityKey],
		//"userName": user.(*model.User).Email,
		"text":     "Hello World.",
	})
}

//TODO railsみたいにpathを返すヘルパーを使いたい
func setupRouter() *gin.Engine {
	router := gin.Default()
	setCors(router)
	authMiddleware := middleware.Auth()

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	v1 := router.Group("/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)
		user := v1.Group("/users")
		{
			user.POST("/create", controller.CreateUser)
		}
		memo := v1.Group("/memos")
		memo.Use(authMiddleware.MiddlewareFunc())
		{
			memo.GET("/index", controller.Index)
			memo.POST("/create", controller.CreateMemo)
		}
		auth := v1.Group("/auth")
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.Use(authMiddleware.MiddlewareFunc())
		{
			auth.GET("/hello", helloHandler)
			//auth.POST("/login", controller.Login)
			auth.POST("/logout", authMiddleware.LogoutHandler)
		}
	}
	return router
}

func setCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"Set-Cookie",
		},
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
}
