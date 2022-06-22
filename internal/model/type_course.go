package model

type CourseType struct {
	CourseTypeID int    `gorm:"primary_key ; AUTO_INCREMENT" json:"-"`
	CourseType   string `json:"coursetype" form:"coursetype"`
}
