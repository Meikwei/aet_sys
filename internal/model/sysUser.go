package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
)

// SysUser 用户表
type SysUser struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	UserName     string `gorm:"column:user_name;type:varchar(50);NOT NULL" json:"userName"`          // 用户名昵称
	UserNumber   string `gorm:"column:user_number;type:varchar(50);NOT NULL" json:"userNumber"`      // 用户编号
	UserPassword string `gorm:"column:user_password;type:varchar(100);NOT NULL" json:"userPassword"` // 用户密码
	UserPhone    string `gorm:"column:user_phone;type:varchar(20)" json:"userPhone"`                 // 用户手机号
	UserAvatar   string `gorm:"column:user_avatar;type:varchar(255)" json:"userAvatar"`              // 用户头像
	CreateUser   string `gorm:"column:create_user;type:varchar(40);NOT NULL" json:"createUser"`      // 创建人
	UpdateUser   string `gorm:"column:update_user;type:varchar(40);NOT NULL" json:"updateUser"`      // 更新人
}

// TableName table name
func (m *SysUser) TableName() string {
	return "sys_user"
}


