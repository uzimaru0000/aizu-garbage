package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/uzimaru0000/aizu-garbage/config"
	"github.com/uzimaru0000/aizu-garbage/controller"
	"github.com/uzimaru0000/aizu-garbage/model"
)

var conf *config.Config
var db *gorm.DB

func main() {
	config.Init("")
	conf = config.Get()
	db = model.DBConnect(conf.MySQL.Host, conf.MySQL)

	engine := gin.Default()

	// 場所情報を保存
	allPlace := controller.GetPlaceList()
	for _, place := range allPlace {
		place.Save(db)
	}

	v1 := engine.Group("/api/v1")
	{
		v1.GET("/places", GetPlaces)
		v1.GET("/garbage/:id", GetGarbageByID)
	}

	engine.Run(":5000")
}

func GetPlaces(c *gin.Context) {
	places := model.GetPlaceAll(db)
	c.JSON(200, gin.H{"places": places})
}

func GetGarbageByID(c *gin.Context) {
	placeID := c.Param("id")

	schedule := &model.Schedule{ID: placeID}
	schedule.Get(db)

	if len(schedule.Categories) == 0 {
		place := &model.Place{PlaceID: placeID}
		place.Get(db)
		var err error
		schedule, err = controller.GetInfo(place)
		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}
		schedule.Save(db)
	}

	c.JSON(200, schedule)
}
