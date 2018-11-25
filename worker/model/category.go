package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	ID      uint32 `gorm:"primary_key"`
	PlaceID string
	Info    string
	Date    time.Time
}

func (c *Category) Save(db *gorm.DB) {
	db.Create(c)
}

func (c *Category) GetCategories(db *gorm.DB) []*Category {
	categories := make([]*Category, 0)

	db.Find(&categories, "place_id = ?", c.PlaceID)

	return categories
}

func (c *Category) Get(db *gorm.DB) {
	db.Find(c, "DATE( date ) = ?", c.Date.Format("2018-06-01"))
}
