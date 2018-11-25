package model_test

import (
	"testing"

	"github.com/uzimaru0000/garbage/worker/model"

	"github.com/uzimaru0000/garbage/worker/config"
)

func TestDBConnect(t *testing.T) {
	config.Init("../../.env")
	conf := config.Get()
	model.DBConnect("127.0.0.1", conf.MySQL)
}
