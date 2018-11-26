package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/uzimaru0000/aizu-garbage/utils/config"
	"github.com/uzimaru0000/aizu-garbage/utils/model"
)

var conf *config.Config
var db *gorm.DB

func main() {
	config.Init("../.env")
	conf = config.Get()
	db = model.DBConnect(conf.MySQL.Host, conf.MySQL)

	engine := gin.Default()

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

	c.JSON(200, schedule)
}
