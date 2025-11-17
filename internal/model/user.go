package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Nickname  string         `gorm:"type:varchar(50)" json:"nickname"`
	Status    int            `gorm:"type:tinyint;default:1" json:"status"` // 1:正常 0:禁用
	Level     int            `gorm:"type:tinyint;default:1" json:"level"` // 1:普通用户 2:VIP用户 3:SVIP用户
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "users"
}

