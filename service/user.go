package service

import (
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
	if err := Db.Create(user).Error; err != nil {
		log.Println(err)
		return  err
	}
	return nil
}

func (UserService) GetUser(email string) (model.User, error) {
	user := model.User{}
	result := Db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user, result.Error
}
