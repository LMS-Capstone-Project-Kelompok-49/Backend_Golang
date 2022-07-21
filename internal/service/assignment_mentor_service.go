package service

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type assignmentMentorService struct {
	repo domain.AssignmentMentorRepository
}

// Delete implements domain.AssignmentMentorService
func (as *assignmentMentorService) Delete(id int) error {
	return as.repo.Delete(id)
}

// Edit implements domain.AssignmentMentorService
func (as *assignmentMentorService) Edit(id int, tugas model.AssignmentMentor) error {
	return as.repo.Update(id, tugas)
}

// GetAllCourses implements domain.AssignmentMentorService
func (as *assignmentMentorService) GetAll() ([]model.AssignmentMentor, error) {
	return as.repo.GetAll()
}

// GetOneCourse implements domain.AssignmentMentorService
func (as *assignmentMentorService) GetOne(id int) (model.AssignmentMentor, error) {
	return as.repo.GetByID(id)
}

// Store implements domain.AssignmentMentorService
func (as *assignmentMentorService) Store(tugas model.AssignmentMentor) error {
	return as.repo.Create(tugas)
}

func NewAssignmentMentorService(repo domain.AssignmentMentorRepository) domain.AssignmentMentorService {
	return &assignmentMentorService{
		repo: repo,
	}
}
