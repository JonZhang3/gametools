package models

import (
	"time"
)

type User struct {
	Id        uint64 `gorm:"primarykey;autoIncrement"`
	Username  string `gorm:"size:64;uniqueIndex;comment:用户名"`
	Nickname  string `gorm:"size:64;comment:昵称"`
	Password  string `gorm:"size:64;comment:密码"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	Id   uint64 `gorm:"primaryKey;autoIncrement"`
	Name string
}
