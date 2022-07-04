package domain

import "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"

type TypeCategoryRepository interface {
	GetAll() []model.CourseType
	GetOne(id int) model.CourseType
}

type TypeCategoryService interface {
	GetAllType() []model.CourseType
	GetOneType(id int) model.CourseType
}
