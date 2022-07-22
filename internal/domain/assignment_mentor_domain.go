package domain

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type AssignmentMentorRepository interface {
	Create(tugas model.AssignmentMentor) error
	Update(id int, tugas model.AssignmentMentor) error
	Delete(id int) error
	GetAll() ([]model.AssignmentMentor, error)
	GetByID(id int) (tugas model.AssignmentMentor, err error)
}

type AssignmentMentorService interface {
	Store(tugas model.AssignmentMentor) error
	Edit(id int, tugas model.AssignmentMentor) error
	Delete(id int) error
	GetAll() ([]model.AssignmentMentor, error)
	GetOne(id int) (model.AssignmentMentor, error)
}
