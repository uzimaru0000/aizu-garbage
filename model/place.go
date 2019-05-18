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
	//if err := db.Find().Error; err != nil { ... }

	return places
}

// gormは多言語のORMを模倣して書かれているのでGOの書き方とは少し離れてはいる
// ので、若干わかりにくいかもしれないけど、エラーハンドリングはした方がいい

// ここの処理もそうだけど、エンティティ（ここでいうPlace)が自身をsaveして、Getしている、またGetPlaceAllは関数では別れているという
// それぞれ、別れてしまっているけど、本来、外部から（今回だとDBから）データを取得するという責務は同じなので、Placeが外部との接続を持たない方がいい
//
//

//
//type Client struct {
//	db *gorm.DB
//}
//

//func NewDBClient(db *gorm.DB) *Client {
//	return &Client{db: db}
//}
//
//func (c *Client) Save(p *Place) error {
//	if err := c.db.Create(p).Error; err != nil {
//		return xerrors.Errorf("place save failed: %w", err)
//	}
//
//	return nil
//}
// interfaceを使うのも拡張性が上がる
