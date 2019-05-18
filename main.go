package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/uzimaru0000/aizu-garbage/config"
	"github.com/uzimaru0000/aizu-garbage/controller"
	"github.com/uzimaru0000/aizu-garbage/model"
)

var (
	conf *config.Config
	// dbクライアントはここではなく、modelの中に隠蔽した方がいい、出ないとmodelを他の場所で生成されていしまう可能性がある
	// もしくは外部とやり取りするクライアント系はまとめてあげてそれからしかmainは呼ばないようにすると良さそう
	db *gorm.DB
)

func main() {
	config.Init("")
	conf = config.Get()

	db = model.DBConnect(conf.MySQL.Host, conf.MySQL)
	defer db.Close() // クローズしてあげたい

	// engine := gin.Default() スコープが無駄に広くなっている

	// 場所情報を保存
	allPlace := controller.GetPlaceList()
	for _, place := range allPlace {
		place.Save(db)
	}
	// 空行
	engine := gin.Default()
	v1 := engine.Group("/api/v1")
	{
		v1.GET("/places", GetPlaces)
		v1.GET("/garbage/:id", GetGarbageByID)
	}

	// error処理
	if err := engine.Run(":5000"); err != nil {
		log.Println(err)
	}
}

// 今の設計上これががmodelとコントローラの二つの役割を持ってしまっているので分離が必要
// 外部で使ってないならプライベートにした方がいい
func GetPlaces(c *gin.Context) {
	places := model.GetPlaceAll(db)
	c.JSON(200, gin.H{"places": places})
}

func GetGarbageByID(c *gin.Context) {
	placeID := c.Param("id")

	schedule := &model.Schedule{ID: placeID}
	schedule.Get(db)

	// ここの判定処理も分けたい
	if len(schedule.Categories) == 0 {
		place := &model.Place{PlaceID: placeID}
		place.Get(db)
		//var err error

		schedule, err := controller.GetInfo(place)
		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}
		// 空行
		schedule.Save(db)
	}

	c.JSON(200, schedule)
}
