package testUtils

// 統合テストする日が来たら役立つだろう

import (
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Mock sqlmock.Sqlmock
var TestDb *gorm.DB

func Before() {
	var err error
	TestDb, Mock, err = getNewDbMock()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	var err error
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return nil, mock, err
	}

	gormDB, err := gorm.Open(mysql.Dialector{
		Config: &mysql.Config{
			DriverName:                "mysql",
			Conn:                      db,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})

	if err != nil {
		panic("miss")
		return gormDB, mock, err
	}

	return gormDB, mock, err
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type AnyString struct{}

func (a AnyString) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}
