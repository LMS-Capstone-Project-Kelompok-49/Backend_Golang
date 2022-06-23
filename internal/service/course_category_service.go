package service

import (
	"net/http"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gopkg.in/go-playground/validator.v9"
)

type courseCategoryService struct {
	repo domain.CourseCategoryRepository
}

// GetAllCourses implements domain.CourseService
func (cs *courseCategoryService) GetAllCategory() []model.CourseCategory {
	return cs.repo.GetAll()
}

// Store implements domain.CourseService
func (cs *courseCategoryService) Store(course model.CourseCategory) (int, error) {
	validate := validator.New()
	err := validate.Struct(course)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, cs.repo.Create(course)
}

func NewCourseCategoryService(repo domain.CourseCategoryRepository) domain.CourseCategoryService {
	return &courseCategoryService{
		repo: repo,
	}
}
