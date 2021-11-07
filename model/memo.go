package model

import (
	"time"
)

//how to comment : https://lunny.gitbooks.io/xorm-manual-en-us/content/chapter-02/4.columns.html

type Memo struct {
	//formのnameと一致 型 xorm:カラム名
	Id int64 `xorm:"pk autoincr int(64)"`
	Text string `xorm:"text" json:"text" form:"text"`
	Discording bool `xorm:"discording  default 0 not null" json:"discording" form:"discording"`
	Permanent bool `xorm:"permanent default 0  not null" json:"permanent" form:"permanent"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
