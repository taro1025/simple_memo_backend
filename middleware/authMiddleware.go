package middleware

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"simple_memo/model"
	"simple_memo/service"
	"time"
)

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginUserJson := session.Get("loginUser")

		if loginUserJson == nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			fmt.Println("miss")
		} else {
			fmt.Println("success")
			c.Next()
		}
	}
}

var identityKey = "id"

type login struct {
	Email    string `form:"email" json:"email" binding:"email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Auth() (*jwt.GinJWTMiddleware) {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				fmt.Println("Payload成功")
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			fmt.Println("Payload失敗")
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			userService := service.UserService{}
			user, result := userService.GetUser(claims[identityKey].(string))
			if !result {
				fmt.Println("Identity失敗")
				return nil
			}
			fmt.Println("Identity成功")
			return &user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Email
			password := loginVals.Password

			userService := service.UserService{}
			user, result := userService.GetUser(email)
			if !result {
				return nil, nil
			}
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				c.Status(http.StatusBadRequest)
			}
			fmt.Println("email" + user.Email)
			return &user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*model.User); ok && v.Email == "rintaro.sakino@msetsu.com" {
				fmt.Println("success" + v.Email)
				return true
			}
			fmt.Println("miss")
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return authMiddleware
}
