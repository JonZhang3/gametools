package models

import "gametools/server/common"

type BaseModel struct {
	ID        uint64          `gorm:"primaryKey;autoIncrement"`
	CreatedAt common.DateTime `gorm:"autoUpdateTime:milli"`
	UpdatedAt common.DateTime `gorm:"autoCreateTime:milli"`
}

type State int8

const (
	Valid   State = 1
	InValid State = 0
	Deleted State = -1
)
