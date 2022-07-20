package repository

import (
	"fmt"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type courseRepoLayer struct {
	DB *gorm.DB
}

// Delete implements domain.CourseRepository
func (cr *courseRepoLayer) Delete(id int) error {

	res := cr.DB.Select("Material").Select("Rating").Select("CourseDetail").Delete(&model.Course{
		CourseID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete data")
	}

	return nil
}

// Update implements domain.CourseRepository
func (cr *courseRepoLayer) Update(id int, course model.Course) error {
	res := cr.DB.Where("course_id = ?", id).UpdateColumns(&course)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update data")
	}
	return nil
}

// GetAll implements domain.CourseRepository
func (cr *courseRepoLayer) GetAll() []model.Course {
	courses := []model.Course{}
	// cr.DB.Find(&courses)
	cr.DB.Preload("User.Profile").Preload("CourseDetail").Preload("Material").Preload("Rating").Find(&courses)

	return courses
}

// GetByID implements domain.CourseRepository
func (cr *courseRepoLayer) GetByID(id int) (course model.Course, err error) {
	res := cr.DB.Where("course_id = ?", id).Preload("User.Profile").Preload("Material").Preload("CourseDetail").Preload("Rating").First(&course)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}
	return
}

// Create implements domain.CourseRepository
func (cr *courseRepoLayer) Create(course model.Course) error {
	res := cr.DB.Create(&course)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create data")
	}
	return nil
}

func NewCourseRepository(db *gorm.DB) domain.CourseRepository {
	return &courseRepoLayer{
		DB: db,
	}
}
