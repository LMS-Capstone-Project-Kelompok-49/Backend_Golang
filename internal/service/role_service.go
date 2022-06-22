package service

import (
	"net/http"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gopkg.in/go-playground/validator.v9"
)

type roleService struct {
	repo domain.RoleRepository
}

// Delete implements domain.RoleService
func (rs *roleService) Delete(id int) error {
	return rs.repo.Delete(id)
}

// Edit implements domain.RoleService
func (rs *roleService) Edit(id int, role model.Role) error {
	return rs.repo.Update(id, role)
}

// GetOneCourse implements domain.RoleService
func (rs *roleService) GetOneRole(id int) (model.Role, error) {
	return rs.repo.GetByID(id)
}

// Store implements domain.RoleService
func (rs *roleService) Store(role model.Role) (int, error) {
	validate := validator.New()
	err := validate.Struct(role)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, rs.repo.Create(role)
}

func NewRoleService(repo domain.RoleRepository) domain.RoleService {
	return &roleService{
		repo: repo,
	}
}
