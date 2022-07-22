package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type ratingService struct {
	repo  domain.RatingRepository
	eRepo domain.EnrollmenRepository
}

// Edit implements domain.RatingService
func (rs *ratingService) Edit(id int, rating model.Rating) (model.Rating, error) {
	return rs.repo.Update(id, rating)
}

// GetByCourse implements domain.RatingService
func (rs *ratingService) GetByCourse(id int) ([]model.Rating, error) {
	return rs.repo.GetAllByCourseID(id)
}

// GetOne implements domain.RatingService
func (rs *ratingService) GetOne(id int) (model.Rating, error) {
	return rs.repo.GetByID(id)
}

// Remove implements domain.RatingService
func (rs *ratingService) Remove(id int) error {
	return rs.repo.Delete(id)
}

// Store implements domain.RatingService
func (rs *ratingService) Store(rating model.Rating) error {
	enrollment := model.Enrollment{}

	id := rating.UserID
	course_id := rating.CourseID

	enr, _ := rs.eRepo.GetByUser(id)

	for i := range enr {
		if enr[i].CourseID == course_id {
			enrollment = enr[i]
		}
	}

	if enrollment.Status != "complete" {
		return fmt.Errorf("selesaikain course dulu")
	}

	log.Println(id)

	data, err := rs.GetByCourse(course_id)
	if err != nil {
		if strings.Contains(err.Error(), "no data") {
			return rs.repo.Create(rating)
		}
		return err
	}

	for i := range data {
		if id == data[i].UserID {
			log.Println(data[i].UserID)
			return fmt.Errorf("satu user satu rating per course")
		}
	}

	return rs.repo.Create(rating)
}

func NewRatingService(repo domain.RatingRepository, eRepo domain.EnrollmenRepository) domain.RatingService {
	return &ratingService{
		repo:  repo,
		eRepo: eRepo,
	}
}
