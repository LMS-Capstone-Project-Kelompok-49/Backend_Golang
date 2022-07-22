package service

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type assignmentUserService struct {
	repo domain.AssignmentUserRepository
}

// GetAllByUserId implements domain.AssignmentUserService
func (as *assignmentUserService) GetAllByUserId(user_id int) ([]model.AssignmentUser, error) {
	return as.repo.GetByUserId(user_id)
}

// GetOne implements domain.AssignmentUserService
func (as *assignmentUserService) GetOne(id int) (model.AssignmentUser, error) {
	return as.repo.GetByID(id)
}

// Edit implements domain.AssignmentUserService
func (as *assignmentUserService) Edit(id int, tugas model.AssignmentUser) (model.AssignmentUser, error) {
	return as.repo.Update(id, tugas)
}

// Store implements domain.AssignmentUserService
func (as *assignmentUserService) Store(tugas model.AssignmentUser) (model.AssignmentUser, error) {
	return as.repo.Create(tugas)
}

func NewAssignmentUserService(repo domain.AssignmentUserRepository) domain.AssignmentUserService {
	return &assignmentUserService{
		repo: repo,
	}
}
