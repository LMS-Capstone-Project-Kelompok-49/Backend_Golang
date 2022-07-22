package model

import (
	"time"

	"gorm.io/gorm"
)

type AssignmentUser struct {
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
	AssignmentUserID   int            `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	AssignmentMentorID int            `json:"-"`
	UserID             int            `json:"-"`
	Document           string         `json:"document" form:"document"`
	Status             string         `json:"-"`
	User               User
}
