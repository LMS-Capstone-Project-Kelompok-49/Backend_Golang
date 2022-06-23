package repository

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type typeCategoryRepoLayer struct {
	DB *gorm.DB
}

// GetAll implements domain.CourseCategoryRepository
func (tcr *typeCategoryRepoLayer) GetAll() []model.CourseType {
	courses := []model.CourseType{}
	tcr.DB.Find(&courses)

	return courses
}

func NewTypeCategoryRepository(db *gorm.DB) domain.TypeCategoryRepository {
	return &typeCategoryRepoLayer{
		DB: db,
	}
}
