package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    int            `gorm:"primary_key; AUTO_INCREMENT" json:"-"`
	Name      string         `json:"name" form:"name"`
	Email     string         `json:"email" form:"email"`
	Password  string         `json:"password" form:"password"`
	RoleID    int            `json:"-"`
	ProfileID int            `json:"-"`
}
