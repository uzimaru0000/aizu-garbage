package model

import (
	"github.com/jinzhu/gorm"
)

type Place struct {
	ID      int    `json:"-"`
	PlaceID string `json:"id"`
	Name    string `json:"name"`
}

func (p *Place) Save(db *gorm.DB) {
	db.Create(p)
}

func (p *Place) Get(db *gorm.DB) {
	db.First(p, "place_id = ? or name = ?", p.PlaceID, p.Name)
}

func GetPlaceAll(db *gorm.DB) []*Place {
	places := []*Place{}
	db.Find(&places)

	return places
}