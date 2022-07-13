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
	Description    string         `json:"description" form:"description"`
	Price          string         `json:"price" form:"price"`
	Rating         string         `json:"rating" form:"rating"`
	Thumbnail      string         `json:"thumbnail" form:"thumbnail"`
}
