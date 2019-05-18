package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/uzimaru0000/aizu-garbage/config"
)

var engine *gorm.DB

func DBConnect(address string, config config.MySQL) *gorm.DB {
	DBMS := "mysql"
	USER := config.User
	PASS := config.Password
	PROTOCOL := "tcp(" + address + ":3306)"
	DBNAME := config.Name
	// CONNECTの生成ロジックを分離した方がいい
	// ここでの責務はデータベースへのコネクションの作成であってそのルールそのものを決めるところではないから
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	// ここではpanicにしない方がいい
	// プロジェクト的には確かにここは最下層に位置しているが、単体で見ればそうではないので、汎用度を上げる意味合いとエラーハンドリングの意味合いで
	// ここで終わらせない方がいい
	if err != nil {
		panic(err.Error())
	}

	return db
}
