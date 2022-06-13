package domain

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type RoleRepository interface {
	Create(role model.Role) error
	Update(id int, role model.Role) error
	Delete(id int) error
	GetByID(id int) (role model.Role, err error)
}

type RoleService interface {
	Store(role model.Role) (int, error)
	Edit(id int, role model.Role) error
	Delete(id int) error
	GetOneRole(id int) (model.Role, error)
}
