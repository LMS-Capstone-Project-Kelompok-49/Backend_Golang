package model

import (
	"time"

	"gorm.io/gorm"
)

type Material struct {
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	MaterialID   int            `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	CourseID     int            `json:"-"`
	MaterialName int            `json:"materialname" form:"materialname"`
	PPT          string         `json:"ppt" form:"ppt"`
	Video        string         `json:"video" form:"video"`
}
