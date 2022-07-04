package domain

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type CourseCategoryRepository interface {
	Create(courseCategory model.CourseCategory) error
	GetAll() []model.CourseCategory
	GetOne(id int) model.CourseCategory
}

type CourseCategoryService interface {
	Store(courseCategory model.CourseCategory) (int, error)
	GetAllCategory() []model.CourseCategory
	GetOneCategory(id int) model.CourseCategory
}
