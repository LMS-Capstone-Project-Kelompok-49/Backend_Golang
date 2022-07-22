package domain

import "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"

type RatingService interface {
	Store(rating model.Rating) error
	Edit(id int, rating model.Rating) (model.Rating, error)
	Remove(id int) error
	GetOne(id int) (model.Rating, error)
	GetByCourse(id int) ([]model.Rating, error)
}

type RatingRepository interface {
	Create(rating model.Rating) error
	Update(id int, rating model.Rating) (model.Rating, error)
	Delete(id int) error
	GetByID(id int) (model.Rating, error)
	GetAllByCourseID(id int) ([]model.Rating, error)
}
