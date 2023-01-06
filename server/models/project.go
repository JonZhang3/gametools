package models

import (
	"time"
)

type Project struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:64;comment:项目名"`
	Desc      string `gorm:"size:256;comment:项目描述"`
	State     State  `gorm:"default:1;comment:状态"`
	CreatedBy uint64 `gorm:"comment:创建用户ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []*User `gorm:"many2many:user_project_rel"`
}
