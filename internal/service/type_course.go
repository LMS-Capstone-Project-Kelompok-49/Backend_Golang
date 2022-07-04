package service

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type typeCourseService struct {
	repo domain.TypeCategoryRepository
}

// GetAllCourses implements domain.CourseService
func (cs *typeCourseService) GetAllType() []model.CourseType {
	return cs.repo.GetAll()
}
func (cs *typeCourseService) GetOneType(id int) model.CourseType {
	return cs.repo.GetOne(id)
}

func NewTypeCategoryService(repo domain.TypeCategoryRepository) domain.TypeCategoryService {
	return &typeCourseService{
		repo: repo,
	}
}
