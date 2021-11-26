package model

import (
"time"
)

//how to comment : https://lunny.gitbooks.io/xorm-manual-en-us/content/chapter-02/4.columns.html

type User struct {
	//formのnameと一致 型 xorm:カラム名
	Id int64 `xorm:"pk autoincr int(64)"`
	Email string `xorm:"email unique" json:"email" form:"email"`
	Password string `xorm:"password" json:"password" form:"password"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
