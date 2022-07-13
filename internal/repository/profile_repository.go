package repository

import (
	"fmt"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type profileRepo struct {
	DB *gorm.DB
}

// Create implements domain.ProfileRepository
func (pr *profileRepo) Create(profile model.UserProfile) error {
	res := pr.DB.Create(&profile)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create data")
	}
	return nil
}

// Delete implements domain.ProfileRepository
func (pr *profileRepo) Delete(id int) error {
	res := pr.DB.Delete(&model.UserProfile{
		ProfileID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete data")
	}

	return nil
}

// GetByID implements domain.ProfileRepository
func (pr *profileRepo) GetByID(id int) (profile model.UserProfile, err error) {
	res := pr.DB.Where("user_id = ?", id).First(&profile)

	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}
	return
}

// Update implements domain.ProfileRepository
func (pr *profileRepo) Update(id int, profile model.UserProfile) error {
	res := pr.DB.Where("user_id = ?", id).UpdateColumns(&profile)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update data")
	}
	return nil
}

func NewProfileRepository(db *gorm.DB) domain.ProfileRepository {
	return &profileRepo{
		DB: db,
	}
}
