package service

import (
	"fmt"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type enrollementService struct {
	repo  domain.CourseRepository
	erepo domain.EnrollmenRepository
}

// GetAllByCourseID implements domain.EnrollmentService
func (*enrollementService) GetAllByCourseID(courseid int) (enr []model.Enrollment, err error) {
	panic("unimplemented")
}

// GetAllByUser implements domain.EnrollmentService
func (*enrollementService) GetAllByUser(id int) ([]model.Enrollment, error) {
	panic("unimplemented")
}

// Join implements domain.EnrollmentService
func (es *enrollementService) Join(enr model.Enrollment, code string) error {
	course, err := es.repo.GetByID(enr.CourseID)
	if err != nil {
		return fmt.Errorf("no course")
	}

	if code != course.Code {
		return fmt.Errorf("kode tidak berlaku")
	}

	return es.erepo.Create(enr)
}

func NewEnrollmentService(erepo domain.EnrollmenRepository, repo domain.CourseRepository) domain.EnrollmentService {
	return &enrollementService{
		erepo: erepo,
		repo:  repo,
	}
}
