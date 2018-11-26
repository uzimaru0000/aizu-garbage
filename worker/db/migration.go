package main

import (
	"os"

	"github.com/uzimaru0000/aizu-garbage/utils/config"
	"github.com/uzimaru0000/aizu-garbage/utils/model"
)

func main() {
	os.Setenv("MODE", "DEV")
	config.Init("../../.env")
	config := config.Get()

	db := model.DBConnect("127.0.0.1", config.MySQL)

	model.Migration(db)
}
