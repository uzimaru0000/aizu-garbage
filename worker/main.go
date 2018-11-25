package main

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/uzimaru0000/garbage/worker/config"
	"github.com/uzimaru0000/garbage/worker/controller"
	"github.com/uzimaru0000/garbage/worker/model"
)

func main() {
	config.Init("../.env")
	conf := config.Get()

	db := model.DBConnect(conf.MySQL.Host, conf.MySQL)

	model.Migration(db)

	getSchedules(db)

	ticker := time.NewTicker(7 * 24 * time.Hour)
	defer ticker.Stop()
	for {
		<-ticker.C
		go getSchedules(db)
	}
}

func getSchedules(db *gorm.DB) {
	places := model.GetPlaceAll(db)
	if len(places) == 0 {
		log.Println("Get place list...")
		places = controller.GetPlaceList()
		for _, p := range places {
			p.Save(db)
		}
	}

	log.Println("Delete schedules...")
	model.DeleteCategories(db)
	log.Println("Get schedules...")
	schedules := controller.GetInfo(places)
	for _, s := range schedules {
		model.SchedulePrint(s)
		s.Save(db)
	}
}