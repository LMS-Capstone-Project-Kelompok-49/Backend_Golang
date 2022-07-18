package model

import (
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CourseID  int            `json:"-"`
	Rating    float32        `json:"rating" form:"rating"`
	Review    string         `json:"review" form:"review"`
}
