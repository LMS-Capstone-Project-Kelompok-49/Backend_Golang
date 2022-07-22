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
	MaterialName string         `json:"materialname" form:"materialname"`
	Description  string         `json:"description" form:"description"`
	PPT          string         `json:"ppt" form:"ppt"`
	Video        string         `json:"video" form:"video"`
}
