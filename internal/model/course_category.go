package model

import (
	"time"

	"gorm.io/gorm"
)

type CourseCategory struct {
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	CourseCategoryID int            `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	Category         string         `json:"category" form:"category"`
}
