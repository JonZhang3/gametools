package models

import (
	"gametools/server/common"
)

type BaseModel struct {
	ID        uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt common.DateTime `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt common.DateTime `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}

type State int8

const (
	Valid   State = 1
	InValid State = -1
	Deleted State = -2
)
