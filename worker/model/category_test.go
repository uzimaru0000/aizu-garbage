package model_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/uzimaru0000/garbage/worker/config"
	"github.com/uzimaru0000/garbage/worker/model"
)

func testingDB() *gorm.DB {
	config.Init("../../.env")
	conf := config.Get()

	db := model.DBConnect("127.0.0.1", conf.MySQL)

	return db
}

// func TestSaveCategory(t *testing.T) {
// 	db := testingDB()

// 	category := model.Category{PlaceID: "000000", Info: "hoge", Date: time.Now().AddDate(0, 0, 1)}

// 	category.Save(db)
// }

func TestGetWeekCategory(t *testing.T) {
	db := testingDB()

	category := model.Category{PlaceID: "000000"}

	cates := category.GetWeekCategory(db)

	for _, c := range cates {
		fmt.Printf("%v\n", *c)
	}
}

func TestGetCategoryByDate(t *testing.T) {
	db := testingDB()

	category := &model.Category{PlaceID: "000000", Date: time.Now().AddDate(0, 0, 1)}

	category.Get(db)

	fmt.Printf("%v\n", *category)
}
