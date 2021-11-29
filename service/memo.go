package service

import (
	"fmt"
	"simple_memo/model"
)

type MemoService struct {}

func (MemoService) SetMemo(user *model.User, memo *model.Memo) error {
	if err := Db.Model(user).Association("Memos").Append(memo); err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

//TODO リミットつけた方がいいかとも思ったけど、1日で消えるメモアプリだしいいか
//err := DbEngine.Limit(10, 0).Find(&memos)  一応こうやると制限つけられる

func (MemoService) Index() []model.Memo{
	memos := make([]model.Memo, 0)
	err := Db.Find(&memos)
	if err != nil {
		panic(err)
	}
	return memos
}
