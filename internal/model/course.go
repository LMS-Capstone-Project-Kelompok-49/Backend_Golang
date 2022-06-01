package model

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	CourseID    int            `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	MentorID    int            `json:"mentorid"`
	CourseType  int            `json:"coursetype" form:"coursetype"`
	Category    int            `json:"category" form:"category"`
	CourseName  string         `json:"coursename" form:"coursename"`
	Description string         `json:"description" form:"description"`
	CoursePrice string         `json:"courseprice" form:"courseprice"`
}
