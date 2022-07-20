package model

import (
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	RatingID  int            `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	CourseID  int            `json:"-"`
	UserID    int            `json:"-"`
	Rating    float64        `json:"rating" form:"rating"`
	Review    string         `json:"review" form:"review"`
}
