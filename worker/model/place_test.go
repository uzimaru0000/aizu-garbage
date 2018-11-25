package model_test

import (
	"testing"

	"github.com/uzimaru0000/garbage/worker/config"
	"github.com/uzimaru0000/garbage/worker/model"
)

func TestGetPlaceByID(t *testing.T) {
	config.Init("../../.env")
	conf := config.Get()

	db := model.DBConnect("127.0.0.1", conf.MySQL)

	place := &model.Place{PlaceID: "000900"}

	place.Get(db)

	if *place != (model.Place{ID: 10, PlaceID: "000900", Name: "上町（A）"}) {
		t.Fatalf("Fatal, geting data is %v\n", *place)
	}
}

func TestGetPlaceByName(t *testing.T) {
	config.Init("../../.env")
	conf := config.Get()

	db := model.DBConnect("127.0.0.1", conf.MySQL)

	place := &model.Place{Name: "上町（A）"}

	place.Get(db)

	if *place != (model.Place{ID: 10, PlaceID: "000900", Name: "上町（A）"}) {
		t.Fatalf("Fatal, geting data is %v\n", *place)
	}
}

func TestGetAllPlace(t *testing.T) {
	config.Init("../../.env")
	conf := config.Get()

	db := model.DBConnect("127.0.0.1", conf.MySQL)

	places := model.GetPlaceAll(db)

	for _, p := range places {
		t.Logf("%v\n", *p)
	}
}
