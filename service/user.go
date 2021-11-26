package service

import (
	"fmt"
	"simple_memo/model"
)

type UserService struct {}

func (UserService) SetUser(user *model.User) error {
	_, err := DbEngine.Insert(user)
	if err!= nil{
		fmt.Println(err)
		return  err
	}
	return nil
}
