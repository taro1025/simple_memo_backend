package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"os"
	"simple_memo/model"
	"simple_memo/testUtils"
	"testing"
)
func TestMain(m *testing.M) {
	testUtils.Before()
	code := m.Run()
	//after()
	os.Exit(code)
}
func TestCreateUser(test *testing.T) {
	testUtils.Mock.ExpectBegin()
	testUtils.Mock.ExpectExec("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`password`) VALUES (?,?,?,?,?)").
		WithArgs(testUtils.AnyTime{},testUtils.AnyTime{},nil,"test@test.com",testUtils.AnyString{}).
		WillReturnResult(sqlmock.NewResult(1, 2))
	testUtils.Mock.ExpectCommit()
	service := UserService{Db: testUtils.TestDb}
	_ = service.SetUser(&model.User{Email: "test@test.com", Password: "password"})
	if err := testUtils.Mock.ExpectationsWereMet(); err != nil {
		test.Errorf("there were unfulfilled expectations: %s", err)
	}
}
