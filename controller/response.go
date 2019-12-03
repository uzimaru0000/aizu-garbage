package controller

import (
	"github.com/uzimaru0000/aizu-garbage/model"
)

type responseSchedule struct {
	Type string `json:"type"`
	Date int64  `json:"date"`
}

type calendar struct {
	Place     *model.Place        `json:"place"`
	Schedules []*responseSchedule `json:"schedules"`
}

func convertSchedule(schedule *model.Schedule) *responseSchedule {
	return &responseSchedule{
		Type: schedule.Type.ToString(),
		Date: schedule.Date.Unix(),
	}
}
