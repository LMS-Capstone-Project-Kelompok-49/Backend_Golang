package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	RoleID      int            `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	RoleName    string         `json:"rolename" form:"rolename"`
	Description string         `json:"description" form:"description"`
}
