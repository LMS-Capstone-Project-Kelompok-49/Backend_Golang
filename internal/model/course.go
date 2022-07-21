package model

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	CourseID     int            `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	CourseName   string         `json:"coursename" form:"coursename"`
	MentorID     int            `json:"-"`
	Code         string         `json:"-"`
	Mentor       User
	CourseDetail CourseDetail
	Student      []Enrollment
	Material     []Material         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Assignment   []AssignmentMentor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating       []Rating           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
