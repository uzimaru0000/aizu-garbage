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

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	return db
}
