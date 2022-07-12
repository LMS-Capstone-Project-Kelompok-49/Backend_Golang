package service

import (
	"net/http"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gopkg.in/go-playground/validator.v9"
)

type profileService struct {
	repo domain.ProfileRepository
}

// Edit implements domain.ProfileService
func (ps *profileService) Edit(user_id int, profile model.UserProfile) error {
	return ps.repo.Update(user_id, profile)
}

// GetProfile implements domain.ProfileService
func (ps *profileService) GetProfile(user_id int) (model.UserProfile, error) {
	return ps.repo.GetByID(user_id)
}

// Remove implements domain.ProfileService
func (ps *profileService) Remove(id int) error {
	return ps.repo.Delete(id)
}

// Store implements domain.ProfileService
func (ps *profileService) Store(profile model.UserProfile) (rescode int, err error) {
	validate := validator.New()
	err = validate.Struct(profile)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = ps.repo.Create(profile)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, err

}

func NewProfileService(repo domain.ProfileRepository) domain.ProfileService {
	return &profileService{
		repo: repo,
	}
}
