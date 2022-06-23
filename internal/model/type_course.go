package model

import (
	"time"

	"gorm.io/gorm"
)

type CourseType struct {
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	CourseTypeID int            `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	CourseType   string         `json:"coursetype" form:"coursetype"`
}
