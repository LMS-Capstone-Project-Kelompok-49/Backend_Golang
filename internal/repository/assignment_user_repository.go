package repository

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type assignmentUserRepo struct {
	DB *gorm.DB
}

// GetByID implements domain.AssignmentUserRepository
func (ar *assignmentUserRepo) GetByID(id int) (model.AssignmentUser, error) {
	tugas := model.AssignmentUser{}
	err := ar.DB.Where("assignment_user_id = ?", id).First(&tugas).Error
	if err != nil {
		return tugas, err
	}
	return tugas, err
}

// GetByUserId implements domain.AssignmentUserRepository
func (ar *assignmentUserRepo) GetByUserId(user_id int) ([]model.AssignmentUser, error) {
	tugass := []model.AssignmentUser{}
	err := ar.DB.Where("user_id = ?", user_id).Find(&tugass).Error
	if err != nil {
		return tugass, err
	}
	return tugass, err
}

// Create implements domain.AssignmentUserRepository
func (ar *assignmentUserRepo) Create(tugas model.AssignmentUser) (model.AssignmentUser, error) {
	err := ar.DB.Create(&tugas).Error
	if err != nil {
		return tugas, err
	}
	return tugas, nil
}

// Update implements domain.AssignmentUserRepository
func (ar *assignmentUserRepo) Update(id int, tugas model.AssignmentUser) (model.AssignmentUser, error) {
	err := ar.DB.Where("assignment_user_id = ?", id).UpdateColumns(&tugas).Error
	if err != nil {
		return tugas, err
	}
	return tugas, nil
}

func NewAssignmentUseerRepository(db *gorm.DB) domain.AssignmentUserRepository {
	return &assignmentUserRepo{
		DB: db,
	}
}
