package model

import (
	"time"

	"gorm.io/gorm"
)

type AdminRole int

const (
	RoleSuperAdmin AdminRole = 1 // 超级管理员
	RoleAdmin      AdminRole = 2 // 普通管理员
	RoleOperator   AdminRole = 3 // 操作员
)

type Admin struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`
	Role      AdminRole      `gorm:"type:tinyint;default:3;not null" json:"role"`
	Status    int            `gorm:"type:tinyint;default:1" json:"status"` // 1:正常 0:禁用
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Admin) TableName() string {
	return "admins"
}

func (r AdminRole) String() string {
	switch r {
	case RoleSuperAdmin:
		return "超级管理员"
	case RoleAdmin:
		return "管理员"
	case RoleOperator:
		return "操作员"
	default:
		return "未知角色"
	}
}

