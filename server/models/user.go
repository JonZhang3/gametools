package models

import (
	"time"
)

type User struct {
	Id        uint64 `gorm:"primarykey"`
	Username  string
	Nickname  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	Id   uint64 `gorm:"primaryKey"`
	Name string
}
