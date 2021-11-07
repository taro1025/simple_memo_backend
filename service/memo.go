package service

import (
	"fmt"
	"simple_memo/model"
)

type MemoService struct {}

func (MemoService) SetMemo(memo *model.Memo) error {
	_, err := DbEngine.Insert(memo)
	if err!= nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (MemoService) Index() []model.Memo{
	tests := make([]model.Memo, 0)
	err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&tests)
	if err != nil {
		panic(err)
	}
	return tests
}