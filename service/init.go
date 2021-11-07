package service

import (
	"database/sql"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"simple_memo/model"
)

var DbEngine *xorm.Engine

func init() {
	user := "root:@tcp(127.0.0.1:3306)/"
	DbName := "simple_memo"
	charSet := "?charset=utf8"
	DSN := user + DbName + charSet

	//DBがなければ作成
	db, err := sql.Open("mysql", user)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + DbName)
	if err != nil {
		panic(err)
	}

	//xormでデータベースと構造体Memoをマッピング
	DbEngine, err = xorm.NewEngine("mysql", DSN)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	DbEngine.ShowSQL(true)           //クエリのログを出す
	DbEngine.SetMaxOpenConns(2)      //同時接続の数？
	DbEngine.Sync2(new(model.Memo)) 	   //NewEngineで指定したsimple_memoDBとMemo構造を同期させる
	fmt.Println("init data base ok")
}
