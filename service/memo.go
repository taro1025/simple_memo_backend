package service

import (
	"fmt"
	"gorm.io/gorm"
	"simple_memo/model"
	"time"
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

func (service *MemoService) Essence(userId uint) []model.Memo{
	memos := make([]model.Memo, 0)
	setDiscord(service, userId)
	err := service.Db.Where("user_id = ? AND discording = ?", userId, false).Find(&memos).Error
	if err != nil {
		return nil
	}
	return memos
}

func (service *MemoService) Discordings(userId uint) []model.Memo{
	memos := make([]model.Memo, 0)
	setDiscord(service, userId)
	err := service.Db.Where("user_id = ? AND discording = ?", userId, true).Find(&memos).Error
	if err != nil {
		return nil
	}
	return memos
}

func (service *MemoService) All() []model.Memo{
	memos := make([]model.Memo, 0)
	err := service.Db.Find(&memos).Error
	if err != nil {
		return nil
	}
	return memos
}

func setDiscord(service *MemoService, userId uint) error {
	discordDate := time.Now().AddDate(0, 0, -3)
	err := service.Db.Model(&model.Memo{}).Where("user_id = ? AND discording = ? AND created_at < ?", userId, false, discordDate).Update("Discording", "1").Error
	if err != nil {
		return err
	}
	return nil
}
