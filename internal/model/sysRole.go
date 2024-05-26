package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
)

// SysRole 角色表
type SysRole struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	RoleName   string `gorm:"column:role_name;type:varchar(50);NOT NULL" json:"roleName"` // 角色名称
	RoleDesc   string `gorm:"column:role_desc;type:varchar(255)" json:"roleDesc"`         // 角色描述
	CreateUser string `gorm:"column:create_user;type:varchar(40)" json:"createUser"`      // 创建人
	UpdateUser string `gorm:"column:update_user;type:varchar(40)" json:"updateUser"`      // 更新人
}

// TableName table name
func (m *SysRole) TableName() string {
	return "sys_role"
}
