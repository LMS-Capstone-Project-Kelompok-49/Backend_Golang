package model

import (
	"time"

	"gorm.io/gorm"
)

type AssignmentMentor struct {
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
	AssignmentMentorID int            `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	CourseID           int            `json:"-"`
	Title              string         `json:"title" form:"title"`
	Refference         string         `json:"refference" form:"refference"`
	Point              int            `json:"point" form:"point"`
}
