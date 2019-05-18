package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Schedule struct {
	ID string `json:"-"`
	//*Place     `json:"place"`
	Place *Place `json:"place"`
	// 確かに変数名を宣言しなくてもGoは大丈夫だけど、定義した方がわかりやすいのと、
	// 展開的な意味合いが出てしまうので書いた方がいい
	Categories []*Category `json:"categories"`
}

// Parserは結構抽象度の高い名前なので、もう少し具体性のある名前を！
func Parser(date string) time.Time {
	now := time.Now()
	// 変数名にもう少し意味合いを持った名前だといい
	a := strings.Split(date, "/")
	b := strings.Split(a[1], "(")

	// error握りつぶしはよくないぜ！
	month, _ := strconv.Atoi(a[0])
	day, _ := strconv.Atoi(b[0])

	// 直接返していいと思うけど、ここもメソッドを分けた方がいい、意味合いがぱっと見わからない
	result := time.Date(now.Year(), time.Month(month), day, 0, 0, 0, 0, now.Location())

	return result
}

// 使ってない？？
func SchedulePrint(schedule *Schedule) {
	fmt.Printf("%s\n", schedule.Place.Name)
	for _, c := range schedule.Categories {
		fmt.Printf("%v %s\n", c.Date, c.Info)
	}
}

// Scheduleのメソッドなのに、実際保存しているのがカテゴリーなのと、複数あるので、メソッド名を変えたほうが良さそう
func (s *Schedule) Save(db *gorm.DB) {
	for _, c := range s.Categories {
		c.PlaceID = s.ID
		c.Save(db)
	}
}

func (s *Schedule) Get(db *gorm.DB) {
	//スコープの広さと処理ブロックをまとめたい
	place := &Place{PlaceID: s.ID}
	place.Get(db)
	s.Place = place

	cate := &Category{PlaceID: s.ID}
	s.Categories = cate.GetCategories(db)
}

// 使ってない？？
func GetAll(db *gorm.DB) []*Schedule {
	result := make([]*Schedule, 0)

	places := GetPlaceAll(db) //空行いらない
	for _, p := range places {
		category := &Category{PlaceID: p.PlaceID}
		result = append(result, &Schedule{ID: p.PlaceID, Place: p, Categories: category.GetCategories(db)})
	}

	return result
}

// 使ってない？？
func DeleteCategories(db *gorm.DB) {
	db.Delete(&Category{})
}

func Migration(db *gorm.DB) {
	place := &Place{}       // ここ変数定義いらないのでは？
	category := &Category{} // ここ変数定義いらないのでは？

	db.CreateTable(place, category)
}
