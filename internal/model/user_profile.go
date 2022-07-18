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
	Fullname       string         `json:"fullname" form:"fullname"`
	DOB            string         `json:"date_of_birth" form:"date_of_birth"`
	ProfilePicture string         `json:"profile_picture" form:"profile_picture"`
	Organization   string         `json:"organization" form:"organization"`
	Job            string         `json:"job" form:"job"`
	Description    string         `json:"description" form:"description"`
}
