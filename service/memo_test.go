package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"simple_memo/testUtils"
	"testing"
	"time"
)

func TestMemoIndex(test *testing.T) {
	//Simulate returned row(s)
	mockRows := sqlmock.NewRows([]string{"created_at"}).AddRow(time.Now())

	testUtils.Mock.ExpectQuery("SELECT * FROM `memos` WHERE `memos`.`deleted_at` IS NULL").
		WillReturnRows(mockRows)

	service := MemoService{Db: testUtils.TestDb}
	_ = service.All()
	if err := testUtils.Mock.ExpectationsWereMet(); err != nil {
		test.Errorf("there were unfulfilled expectations: %s", err)
	}
}
