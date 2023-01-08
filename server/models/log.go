package models

// OperationLog 操作日志
type OperationLog struct {
	BaseModel
	Type    OperationType `gorm:"not null;comment:操作类型"`
	Content string        `gorm:"type:text;comment:操作内容"`
	UserId  uint64        `gorm:"not null"`
}
