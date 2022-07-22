package repository

import (
	"fmt"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type ratingRepository struct {
	DB *gorm.DB
}

// Create implements domain.RatingRepository
func (rr *ratingRepository) Create(rating model.Rating) error {
	res := rr.DB.Create(&rating)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create data")
	}
	return nil
}

// Delete implements domain.RatingRepository
func (rr *ratingRepository) Delete(id int) error {
	res := rr.DB.Delete(&model.Rating{
		RatingID: id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete data")
	}
	return nil
}

// GetAllByCourseID implements domain.RatingRepository
func (rr *ratingRepository) GetAllByCourseID(id int) ([]model.Rating, error) {
	ratings := []model.Rating{}
	res := rr.DB.Where("course_id = ?", id).Find(&ratings)

	if res.RowsAffected < 1 {
		return ratings, fmt.Errorf("no data")
	}
	return ratings, nil
}

// GetByID implements domain.RatingRepository
func (rr *ratingRepository) GetByID(id int) (model.Rating, error) {
	rating := model.Rating{}
	res := rr.DB.Where("rating_id = ?", id).First(&rating)
	if res.RowsAffected < 1 {
		return rating, fmt.Errorf("not found")
	}
	return rating, nil
}

// Update implements domain.RatingRepository
func (rr *ratingRepository) Update(id int, rating model.Rating) (model.Rating, error) {
	res := rr.DB.Where("rating_id = ?", id).UpdateColumns(&rating)
	if res.RowsAffected < 1 {
		return rating, fmt.Errorf("error updae data")
	}
	return rating, nil
}

func NewRatingRepository(db *gorm.DB) domain.RatingRepository {
	return &ratingRepository{
		DB: db,
	}
}
