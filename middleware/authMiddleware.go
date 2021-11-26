package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
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
