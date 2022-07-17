package model

import (
	"time"

	"gorm.io/gorm"
)

type CourseDetail struct {
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	CourseDetailID int            `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	CourseID       int
	CategoryID     int            `json:"category_id" form:"category_id"`
	Description    string         `json:"description" form:"description"`
	Price          string         `json:"price" form:"price"`
	Rating         string         `json:"rating" form:"rating"`
	Avatar         string         `json:"thumbnail" form:"thumbnail"`
	Media          string         `json:"media" form:"media"`
}
