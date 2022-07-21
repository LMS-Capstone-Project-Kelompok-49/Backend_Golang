package repository

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type enrollmentRepoLayer struct {
	DB *gorm.DB
}

// Create implements domain.EnrollmenRepository
func (er *enrollmentRepoLayer) Create(enr model.Enrollment) error {
	err := er.DB.Create(&enr).Error
	if err != nil {
		return err
	}
	return nil
}

// GetByCourse implements domain.EnrollmenRepository
func (*enrollmentRepoLayer) GetByCourse(course_id int) (enr []model.Enrollment, err error) {
	panic("unimplemented")
}

// GetByUser implements domain.EnrollmenRepository
func (*enrollmentRepoLayer) GetByUser(user_id int) (enr model.Enrollment, err error) {
	panic("unimplemented")
}

func NewEnrollmentRepository(db *gorm.DB) domain.EnrollmenRepository {
	return &enrollmentRepoLayer{
		DB: db,
	}
}
