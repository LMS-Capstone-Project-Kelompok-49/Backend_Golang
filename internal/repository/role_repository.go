package repository

import (
	"fmt"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type roleRepoLayer struct {
	DB *gorm.DB
}

// Create implements domain.RoleRepository
func (rr *roleRepoLayer) Create(role model.Role) error {
	res := rr.DB.Create(&role)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create data")
	}
	return nil
}

// Delete implements domain.RoleRepository
func (rr *roleRepoLayer) Delete(id int) error {
	res := rr.DB.Delete(&model.Role{
		RoleID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete data")
	}

	return nil
}

// GetByID implements domain.RoleRepository
func (rr *roleRepoLayer) GetByID(id int) (role model.Role, err error) {
	res := rr.DB.Where("course_id = ?", id).First(&role)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}
	return
}

// Update implements domain.RoleRepository
func (rr *roleRepoLayer) Update(id int, role model.Role) error {
	res := rr.DB.Where("role_id = ?", id).UpdateColumns(&role)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update data")
	}
	return nil
}

func NewRoleRepository(db *gorm.DB) domain.RoleRepository {
	return &roleRepoLayer{
		DB: db,
	}
}
