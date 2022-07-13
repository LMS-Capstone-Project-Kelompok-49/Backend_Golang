package model

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	CourseID   int            `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	MentorID   int            `json:"-"`
	User       User           `gorm:"foreignKey:MentorID"`
	CourseType string         `json:"coursetype" form:"coursetype"`
	Category   string         `json:"category" form:"category"`
	CourseName string         `json:"coursename" form:"coursename"`
}
