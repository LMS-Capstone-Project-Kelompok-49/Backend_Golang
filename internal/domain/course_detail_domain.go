package domain

import "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"

type CourseDetailRepository interface {
	Update(id int, detail model.CourseDetail) error
	GetByID(id int) (course model.CourseDetail, err error)
}

type CourseDetailService interface {
	Edit(id int, detail model.CourseDetail) error
	GetOne(id int) (model.CourseDetail, error)
}
