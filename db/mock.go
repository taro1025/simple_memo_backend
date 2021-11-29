package db

//func GetNewDbMock() (*sql.DB, sqlmock.Sqlmock, error) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		return nil, mock, err
//	}
//
//	gormDB, err := xorm.NewEngine("mysql", service.DSN)
//
//	if err != nil {
//		return gormDB, mock, err
//	}
//
//	return gormDB, mock, err
//}