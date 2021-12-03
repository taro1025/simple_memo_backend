package service

import (
	"fmt"
	"gorm.io/gorm"
	"simple_memo/model"
)

type MemoService struct {
	Db *gorm.DB
}

func (service *MemoService) SetMemo(user *model.User, memo *model.Memo) error {
	if err := service.Db.Model(user).Association("Memos").Append(memo); err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

//TODO リミットつけた方がいいかとも思ったけど、1日で消えるメモアプリだしいいか
//err := DbEngine.Limit(10, 0).Find(&memos)  一応こうやると制限つけられる

func (service *MemoService) Index() []model.Memo{
	memos := make([]model.Memo, 0)
	err := service.Db.Find(&memos).Error
	if err != nil {
		//panic(err)
		return nil
	}
	return memos
}
