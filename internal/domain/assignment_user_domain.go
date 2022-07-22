package domain

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type AssignmentUserRepository interface {
	Create(tugas model.AssignmentUser) (model.AssignmentUser, error)
	Update(id int, tugas model.AssignmentUser) (model.AssignmentUser, error)
	GetByID(id int) (model.AssignmentUser, error)
	GetByUserId(user_id int) ([]model.AssignmentUser, error)
}

type AssignmentUserService interface {
	Store(tugas model.AssignmentUser) (model.AssignmentUser, error)
	Edit(id int, tugas model.AssignmentUser) (model.AssignmentUser, error)
	GetOne(id int) (model.AssignmentUser, error)
	GetAllByUserId(user_id int) ([]model.AssignmentUser, error)
}
