package service

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type courseDetailService struct {
	repo domain.CourseDetailRepository
}

// Edit implements domain.CourseDetailService
func (cd *courseDetailService) Edit(id int, detail model.CourseDetail) error {
	return cd.repo.Update(id, detail)
}

// GetOne implements domain.CourseDetailService
func (cd *courseDetailService) GetOne(id int) (model.CourseDetail, error) {
	return cd.repo.GetByID(id)
}

func NewCourseDetailService(repo domain.CourseDetailRepository) domain.CourseDetailService {
	return &courseDetailService{
		repo: repo,
	}
}
