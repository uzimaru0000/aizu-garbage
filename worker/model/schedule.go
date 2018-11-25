package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Schedule struct {
	ID         string
	Place      *Place
	Categories []*Category
}

func Parser(date string) time.Time {
	now := time.Now()
	a := strings.Split(date, "/")
	b := strings.Split(a[1], "(")

	month, _ := strconv.Atoi(a[0])
	day, _ := strconv.Atoi(b[0])

	result := time.Date(now.Year(), time.Month(month), day, 0, 0, 0, 0, now.Location())

	return result
}

func SchedulePrint(schedule *Schedule) {
	fmt.Printf("%s\n", schedule.Place.Name)
	for _, c := range schedule.Categories {
		fmt.Printf("%v %s\n", c.Date, c.Info)
	}
}

func (s *Schedule) Save(db *gorm.DB) {
	for _, c := range s.Categories {
		c.PlaceID = s.ID
		c.Save(db)
	}
}

func (s *Schedule) Get(db *gorm.DB) {
	place := &Place{PlaceID: s.ID}
	cate := &Category{PlaceID: s.ID}
	place.Get(db)

	s.Place = place
	s.Categories = cate.GetCategories(db)
}

func GetAll(db *gorm.DB) []*Schedule {
	result := make([]*Schedule, 0)

	places := GetPlaceAll(db)

	for _, p := range places {
		category := &Category{PlaceID: p.PlaceID}
		result = append(result, &Schedule{ID: p.PlaceID, Place: p, Categories: category.GetCategories(db)})
	}

	return result
}

func DeleteCategories(db *gorm.DB) {
	db.Delete(&Category{})
}

func Migration(db *gorm.DB) {
	place := &Place{}
	category := &Category{}

	db.CreateTable(place, category)
}
