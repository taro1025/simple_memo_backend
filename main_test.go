package main

import (
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"simple_memo/service"
	"testing"
	"time"
)

var mock sqlmock.Sqlmock
var db *gorm.DB

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	//after()
	os.Exit(code)
}

func TestLogin(test *testing.T){
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/login", nil)

}
//
//func TestCreateUser(test *testing.T) {
//	mock.ExpectBegin()
//	mock.ExpectExec("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`password`) VALUES (?,?,?,?,?)").
//		WithArgs(AnyTime{},AnyTime{},nil,"test@test.com",AnyString{}).
//		WillReturnResult(sqlmock.NewResult(1, 2))
//	mock.ExpectCommit()
//	service := service.UserService{Db: db}
//	_ = service.SetUser(&model.User{Email: "test@test.com", Password: "password"})
//	if err := mock.ExpectationsWereMet(); err != nil {
//		test.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}

func TestMemoIndex(test *testing.T) {
	//Simulate returned row(s)
	mockRows := sqlmock.NewRows([]string{"created_at"}).AddRow(time.Now())

	mock.ExpectQuery("SELECT * FROM `memos` WHERE `memos`.`deleted_at` IS NULL").
		WillReturnRows(mockRows)

	service := service.MemoService{Db: db}
	_ = service.Index()
	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func before() {
	var err error
	db, mock, err = GetNewDbMock()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
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