package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_memo/model"
	"simple_memo/service"

	//jwtmiddleware "github.com/auth0/go-jwt-middleware"
	//jwt "github.com/form3tech-oss/jwt-go"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("loginUser", c.PostForm("userId"))
	session.Save()
	c.String(http.StatusOK, "ログイン完了")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.String(http.StatusOK, "ログアウトしました")
}

func CreateUser(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil{
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	userService :=service.UserService{}
	err = userService.SetUser(&user)
	if err != nil{
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

//
//func GetToken(c *gin.Context) {
//
//	// headerのセット
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	// claimsのセット
//	claims := token.Claims.(jwt.MapClaims)
//	claims["admin"] = true
//	claims["sub"] = "54546557354"
//	claims["name"] = "taro"
//	claims["iat"] = time.Now()
//	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
//
//	// 電子署名
//	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
//
//	// JWTを返却
//	c.Write([]byte(tokenString))
//}
//
//// JwtMiddleware check token
//var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
//	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
//		return []byte(os.Getenv("SIGNINGKEY")), nil
//	},
//	SigningMethod: jwt.SigningMethodHS256,
//})