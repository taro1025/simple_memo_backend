package service

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"simple_memo/model"
)

type UserService struct {
	Db *gorm.DB
}

func (service *UserService) SetUser(user *model.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hash)
	if err := service.Db.Create(user).Error; err != nil {	//service.Db.Debug().Createでsqlを表示できる
		log.Println(err)
		return  err
	}
	return nil
}

func (UserService) GetUser(email string) (model.User, error) {
	user := model.User{}
	result := Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user, result.Error
}
