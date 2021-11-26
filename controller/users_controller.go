package controller

//func CreateUser(c *gin.Context) {
//	user := model.User{}
//	err := c.Bind(&user)
//	if err != nil{
//		c.String(http.StatusBadRequest, "Bad request")
//		return
//	}
//	userService :=service.UserService{}
//	err = userService.SetUser(&user)
//	if err != nil{
//		c.String(http.StatusInternalServerError, "Server Error")
//		return
//	}
//	c.JSON(http.StatusCreated, gin.H{
//		"status": "ok",
//	})
//}
