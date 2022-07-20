package service

import (
	"net/http"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gopkg.in/go-playground/validator.v9"
)

type courseService struct {
	repo domain.CourseRepository
}

// Delete implements domain.CourseService
func (cs *courseService) Delete(id int) error {
	return cs.repo.Delete(id)
}

// Edit implements domain.CourseService
func (cs *courseService) Edit(id int, course model.Course) error {
	return cs.repo.Update(id, course)
}

// GetAllCourses implements domain.CourseService
func (cs *courseService) GetAllCourses() []model.Course {
	res := cs.repo.GetAll()

	return res
}

// GetOneCourse implements domain.CourseService
func (cs *courseService) GetOneCourse(id int) (model.Course, error) {
	return cs.repo.GetByID(id)
}

// Store implements domain.CourseService
func (cs *courseService) Store(course model.Course) (int, error) {
	validate := validator.New()
	err := validate.Struct(course)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, cs.repo.Create(course)
}

func NewCourseService(repo domain.CourseRepository) domain.CourseService {
	return &courseService{
		repo: repo,
	}
}
