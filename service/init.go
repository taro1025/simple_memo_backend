package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"simple_memo/model"
)

var Db *gorm.DB
func init() {
	user := "root:@tcp(127.0.0.1:3306)/"
	DbName := "simple_memo"
	charSet := "?charset=utf8&parseTime=True"
	DSN := user + DbName + charSet

	var err error
	Db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Db.AutoMigrate(&model.User{}, &model.Memo{})
	Db.Logger.LogMode(logger.Info)
	log.Println("init data base ok")
}
