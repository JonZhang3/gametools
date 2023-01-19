package models

import (
	"gametools/server/app"
	"gametools/server/common"
	"gametools/server/tools"
	"strings"
)

type Project struct {
	BaseModel
	Name        string  `gorm:"unique;not null;size:64;comment:项目名" form:"name" json:"name"`
	Description string  `gorm:"size:256;comment:项目描述" form:"description" json:"description"`
	State       State   `gorm:"default:1;comment:状态" json:"state"`
	CreatedBy   uint64  `gorm:"not null;comment:创建用户ID" json:"createdBy"`
	Users       []*User `gorm:"many2many:user_project_rel" json:"users"`
}

func PageListProjects(proj *Project, offset, limit int) *app.Pager {
	tx := app.DB.Model(&Project{}).
		Where("state = ?", tools.If(proj.State == 0, Valid, proj.State))
	if proj.Name != "" {
		tx.Where("lower(name) like ?", "%"+strings.ToLower(proj.Name)+"%")
	}
	var count int64
	tx.Count(&count)
	var list []Project = nil
	tx.Offset(offset).Limit(limit).Order("created_at desc").Find(&list)
	return &app.Pager{
		Total: count,
		Data:  list,
	}
}

func FindProjectByName(name string) *Project {
	p := &Project{}
	result := app.DB.Where("name = ?", name).Take(p)
	if result.RowsAffected == 0 {
		return nil
	}
	return p
}

func ChangeProjectState(id uint64, state State) {
	app.DB.Model(&Project{}).
		Where("id = ?", id).
		Update("state", state)
}

type DocumentType string

const (
	DocumentTypeCSV  = "csv"
	DocumentTypeXML  = "xml"
	DocumentTypeJSON = "json"
)

type Document struct {
	BaseModel
	Name        string       `gorm:"not null;size:256;uniqueIndex:idx_name_project;comment:文档名称"`
	Description string       `gorm:"size:256;comment:文档描述"`
	Type        DocumentType `gorm:"default:csv;comment:文档类型"`
	StoragePath string       `gorm:"comment:存储路径"`
	ProjectId   uint64       `gorm:"uniqueIndex:idx_name_project;comment:项目ID"`
	Sort        uint32       `gorm:"default:0;comment:显示顺序"`
	CreatedBy   uint64       `gorm:"not null;comment:创建用户"`
}

type ColumnType string

const (
	ColumnTypeText   = "text"
	ColumnTypeNumber = "number"
)

type ColumnCalcType string

const (
	ColumnCalcTypeFixed = "fixed"
	ColumnCalcTypeEdit  = "edit"
	ColumnCalcTypeCalc  = "calc"
)

type DocumentColumn struct {
	BaseModel
	Name        string         `gorm:"not null;size:256;comment:文档列名"`
	Description string         `gorm:"not null;size:256;comment:列描述"`
	Type        ColumnType     `gorm:"not null;size:256;comment:列类型"`
	CalcType    ColumnCalcType `gorm:"default:edit;not null;size:64;comment:计算方式"`
	Size        uint64         `gorm:"comment:最大长度/大小"`
	Sort        uint32         `gorm:"default:0;comment:显示顺序"`
	CreatedBy   uint64         `gorm:"not null;comment:创建用户"`
}

type CommitLog struct {
	CommitHash string
	Committer  string
	Email      string
	Message    string
	Date       common.DateTime
}
