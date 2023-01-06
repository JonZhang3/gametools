package models

import (
	"time"
)

type User struct {
	ID        uint64 `gorm:"primarykey;autoIncrement"`
	Username  string `gorm:"size:64;uniqueIndex;comment:用户名" form:"username"`
	Nickname  string `gorm:"size:64;comment:昵称" form:"nickname"`
	Password  string `gorm:"size:64;comment:密码" form:"password"`
	State     State  `gorm:"default:1;comment:状态"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Projects  []*Project `gorm:"many2many:user_project_rel"`
}

type Role struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Name      string
	State     State `gorm:"default:1;comment:状态"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
