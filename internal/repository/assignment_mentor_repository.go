package repository

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type assignmentMentorRepo struct {
	DB *gorm.DB
}

// Create implements domain.AssignmentMentorRepository
func (ar *assignmentMentorRepo) Create(tugas model.AssignmentMentor) error {
	err := ar.DB.Create(&tugas).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete implements domain.AssignmentMentorRepository
func (ar *assignmentMentorRepo) Delete(id int) error {
	err := ar.DB.Delete(&model.AssignmentMentor{
		AssignmentMentorID: id,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements domain.AssignmentMentorRepository
func (ar *assignmentMentorRepo) GetAll() ([]model.AssignmentMentor, error) {
	tugass := []model.AssignmentMentor{}
	err := ar.DB.Find(&tugass).Error
	if err != nil {
		return tugass, err
	}
	return tugass, nil
}

// GetByID implements domain.AssignmentMentorRepository
func (ar *assignmentMentorRepo) GetByID(id int) (tugas model.AssignmentMentor, err error) {
	err = ar.DB.Where("assignment_mentor_id = ?", id).Preload("AssignmentUser").First(&tugas).Error
	if err != nil {
		return tugas, err
	}
	return tugas, nil
}

// Update implements domain.AssignmentMentorRepository
func (ar *assignmentMentorRepo) Update(id int, tugas model.AssignmentMentor) error {
	err := ar.DB.Where("assignment_mentor_id = ?", id).UpdateColumns(&tugas).Error
	if err != nil {
		return err
	}
	return nil
}

func NewAssignmentMentorRepository(db *gorm.DB) domain.AssignmentMentorRepository {
	return &assignmentMentorRepo{
		DB: db,
	}
}
