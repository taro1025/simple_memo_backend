package model

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	UserID int
	Text string `gorm:"type: text"`
	Discording bool `gorm:"default: 0; not null"`
	Permanent bool `xorm:"default: 0; not null"`
}
