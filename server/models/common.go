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

type ResourceType string

const (
	ResourceTypeMenu ResourceType = "menu"
	ResourceTypePath ResourceType = "path"
)

type DocumentType string

const (
	DocumentTypeCSV  = "csv"
	DocumentTypeXML  = "xml"
	DocumentTypeJSON = "json"
)

type ColumnCalcType string

const (
	ColumnCalcTypeEdit = "edit"
	ColumnCalcTypeCalc = "calc"
)

type OperationType uint8

const (
	OperationLogin = iota + 1
	OperationCreateProject
	OperationUpdateProject
	OperationDeleteProject
	OperationAddUser
	OperationDeleteUser
	OperationCreateDocument
	OperationDeleteDocument
)
