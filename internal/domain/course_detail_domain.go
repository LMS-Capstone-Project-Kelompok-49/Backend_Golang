package domain

import "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"

type CourseDetailRepository interface {
	Update(id int, detail model.CourseDetail) (res model.CourseDetail, err error)
	GetByID(id int) (detail model.CourseDetail, err error)
}

type CourseDetailService interface {
	Edit(id int, detail model.CourseDetail) (res model.CourseDetail, err error)
	GetOne(id int) (model.CourseDetail, error)
}
