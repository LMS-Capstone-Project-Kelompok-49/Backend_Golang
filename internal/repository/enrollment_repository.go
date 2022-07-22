package repository

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type enrollmentRepoLayer struct {
	DB *gorm.DB
}

// UpdateStatus implements domain.EnrollmenRepository
func (er *enrollmentRepoLayer) UpdateStatus(user_id int, course_id int, enr model.Enrollment) error {
	err := er.DB.Model(&enr).Where("user_id = ?", user_id).Where("course_id = ?", course_id).Update("status", "complete").Update("progress", gorm.Expr("progress + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

// Update implements domain.EnrollmenRepository
func (er *enrollmentRepoLayer) Update(user_id int, course_id int, enr model.Enrollment) error {
	err := er.DB.Model(&enr).Where("user_id = ?", user_id).Where("course_id = ?", course_id).Update("progress", gorm.Expr("progress + ?", 1)).Update("status", "ongoing").Error
	if err != nil {
		return err
	}
	return nil
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
func (er *enrollmentRepoLayer) GetByUser(user_id int) (enr []model.Enrollment, err error) {
	err = er.DB.Where("user_id = ?", user_id).Find(&enr).Error
	if err != nil {
		return enr, err
	}
	return enr, err
}

func NewEnrollmentRepository(db *gorm.DB) domain.EnrollmenRepository {
	return &enrollmentRepoLayer{
		DB: db,
	}
}
