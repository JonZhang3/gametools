package models

type User struct {
	BaseModel
	Username string     `gorm:"not null;size:64;unique;comment:用户名" form:"username"`
	Nickname string     `gorm:"size:64;comment:昵称" form:"nickname"`
	Password string     `gorm:"not null;size:64;comment:密码" form:"password"`
	Email    string     `gorm:"comment:邮箱" form:"email"`
	IsAdmin  bool       `gorm:"default:false;comment:是否是管理员"`
	State    State      `gorm:"default:1;comment:状态"`
	Projects []*Project `gorm:"many2many:user_project_rel"`
}

type Role struct {
	BaseModel
	Name        string `gorm:"not null;size:64;comment:角色名称"`
	Code        string `gorm:"not null;unique;size:64;comment:角色唯一标识"`
	Description string `gorm:"size:256;comment:描述信息"`
	State       State  `gorm:"default:1;comment:状态"`
}

type Resource struct {
	BaseModel
	Name        string       `gorm:"not null;size:64;comment:资源名称"`
	Description string       `gorm:"size:256;comment:描述信息"`
	Type        ResourceType `gorm:"not null;comment:资源类型"`
	Permission  string       `gorm:"not null;unique;size:256;comment:权限标识"`
	Path        string       `gorm:"size:256;comment:路由地址"`
	State       State        `gorm:"default:1;comment:状态"`
}

type UserRole struct {
	UserId    uint64 `gorm:"not null"`
	RoleId    uint64 `gorm:"not null"`
	ProjectId uint64 `gorm:"not null"`
}

func (UserRole) TableName() string {
	return "user_role_rel"
}
