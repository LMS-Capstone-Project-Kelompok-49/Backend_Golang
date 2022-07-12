package model

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	ProfileID      int            `gorm:"primary_key; AUTO_INCREMENT" json:"-"`
	UserID         int            `json:"-"`
	Fullname       string         `json:"fullname" form:"fullname" validate:"required"`
	DOB            string         `json:"dateofbirth" form:"dateofbirth" validate:"required"`
	ProfilePicture string         `json:"profilepicture" form:"profilepicture"`
	Organization   string         `json:"organization" form:"organization"`
}
