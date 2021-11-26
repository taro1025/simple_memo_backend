package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"simple_memo/model"
)

type UserService struct {}

func (UserService) SetUser(user *model.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hash)
	_, err = DbEngine.Insert(user)
	if err!= nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (UserService) GetUser(email string) (model.User, bool) {
	user := model.User{}
	result, err := DbEngine.Where("email = ?", email).Get(&user)
	if err != nil {
		panic(err)
	}
	if !result {
		log.Fatal("Not Found")
	}
	return user, result
}
