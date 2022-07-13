package domain

import "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"

type ProfileRepository interface {
	Create(profile model.UserProfile) error
	GetByID(user_id int) (profile model.UserProfile, err error)
	Delete(id int) error
	Update(user_id int, profile model.UserProfile) error
}

type ProfileService interface {
	Store(profile model.UserProfile) (profile_id int, err error)
	Edit(user_id int, profile model.UserProfile) error
	Remove(id int) error
	GetProfile(user_id int) (model.UserProfile, error)
}
