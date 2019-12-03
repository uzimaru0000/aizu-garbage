package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/uzimaru0000/aizu-garbage/client"
	"github.com/uzimaru0000/aizu-garbage/config"
	"github.com/uzimaru0000/aizu-garbage/controller"
	"github.com/uzimaru0000/aizu-garbage/repository"
	"github.com/uzimaru0000/aizu-garbage/usecase"
)

func main() {
	mySQL := &config.MySQL{
		User:     os.Getenv("DB_USER"),
		PassWord: os.Getenv("DB_PASS"),
		Address:  os.Getenv("DB_ADDRESS"),
		Port:     os.Getenv("DB_PORT"),
		Socket:   os.Getenv("DB_SOCKET"),
		DataBase: os.Getenv("DB_NAME"),
	}
	db, err := sql.Open("mysql", mySQL.Source())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	scheduleRepo := repository.NewScheduleRepository(db)
	placeRepo := repository.NewPlaceRepository(db)

	fetcher := client.NewClient(http.DefaultClient)

	scheduleCase := usecase.NewScheduleUseCase(scheduleRepo, fetcher)
	placeCase := usecase.NewPlaceUseCase(placeRepo)

	controller := controller.NewController(scheduleCase, placeCase, os.Getenv("DELETE_KEY"))

	controller.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
