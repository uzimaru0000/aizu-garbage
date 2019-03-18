package main

import (
	"log"

	"github.com/uzimaru0000/aizu-garbage/config"
	"github.com/uzimaru0000/aizu-garbage/model"
)

func main() {
	config.Init("")
	config := config.Get()

	db := model.DBConnect(config.MySQL.Host, config.MySQL)

	model.Migration(db)
	log.Print("Success migrate!")
}
