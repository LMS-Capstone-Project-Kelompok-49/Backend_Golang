package domain

import "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"

type TypeCategoryRepository interface {
	GetAll() []model.CourseType
}

type TypeCategoryService interface {
	GetAllType() []model.CourseType
}
