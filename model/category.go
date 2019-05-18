package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	ID      uint32    `gorm:"primary_key" json:"-"`
	PlaceID string    `json:"-"`
	Info    string    `json:"info"`
	Date    time.Time `json:"date"`
}

func (c *Category) Save(db *gorm.DB) {
	db.Create(c)
}

func (c *Category) GetCategories(db *gorm.DB) []*Category {
	categories := make([]*Category, 0) // 空行いらない
	db.Find(&categories, "place_id = ?", c.PlaceID)

	return categories
}

func (c *Category) Get(db *gorm.DB) {
	// c.Date.Format("2018-06-01"))は別関数で切り出した方がいい
	// ここにこのフォーマットでデータを加工するところまでが表現されてしまっているので、外部で表現して、それを呼び出した方が
	// 変更も容易だし、読みやすい
	db.Find(c, "DATE( date ) = ?", c.Date.Format("2018-06-01"))
}

// Placeと同じことが言える
