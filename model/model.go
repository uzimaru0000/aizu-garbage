package model

import "time"

type ScheduleType int

const (
	None ScheduleType = iota
	Barnable
	UnBarnable
	Can
	Bottle
	Paper
	PET
	Plastic
	Oversize
)

// 地区
type Place struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// ゴミの種類
type Schedule struct {
	Type      ScheduleType `db:"type"`
	Date      time.Time    `db:"date"`
	CreatedAt time.Time    `db:"created_at"`
}

// 文字列からIDに変換
func ScheduleTypeFromString(str string) ScheduleType {
	switch str {
	case "燃やせるごみ":
		return Barnable
	case "燃やせないごみ":
		return UnBarnable
	case "かん類":
		return Can
	case "びん類":
		return Bottle
	case "古紙類":
		return Paper
	case "ペットボトル":
		return PET
	case "プラスチック製容器包装":
		return Plastic
	case "粗大ごみ":
		return Oversize
	}
	return None
}

func (t ScheduleType) ToString() string {
	switch t {
	case Barnable:
		return "燃えるごみ"
	case UnBarnable:
		return "燃えないごみ"
	case Can:
		return "かん類"
	case Bottle:
		return "びん類"
	case Paper:
		return "古紙類"
	case PET:
		return "ペットボトル"
	case Plastic:
		return "プラスチック製容器包装"
	case Oversize:
		return "粗大ごみ"
	}

	return "なし"
}
