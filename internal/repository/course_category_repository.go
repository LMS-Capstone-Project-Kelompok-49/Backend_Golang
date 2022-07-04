package repository

import (
	"fmt"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type courseCategoryRepoLayer struct {
	DB *gorm.DB
}

// GetAll implements domain.CourseCategoryRepository
func (ccr *courseCategoryRepoLayer) GetAll() []model.CourseCategory {
	courses := []model.CourseCategory{}
	ccr.DB.Find(&courses)

	return courses
}

func (ccr *courseCategoryRepoLayer) GetOne(id int) model.CourseCategory {
	courses := model.CourseCategory{}
	ccr.DB.Where("course_category_id = ?", id).Find(&courses)

	return courses
}

// Create implements domain.CourseCategoryRepository
func (ccr *courseCategoryRepoLayer) Create(course model.CourseCategory) error {
	res := ccr.DB.Create(&course)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create data")
	}
	return nil
}

func NewCourseCategoryRepository(db *gorm.DB) domain.CourseCategoryRepository {
	return &courseCategoryRepoLayer{
		DB: db,
	}
}
