package repository

import (
	"fmt"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type detailRepo struct {
	DB *gorm.DB
}

// GetByID implements domain.CourseDetailRepository
func (dr *detailRepo) GetByID(id int) (course model.CourseDetail, err error) {
	detail := model.CourseDetail{}
	res := dr.DB.Where("course_id = ?", id).First(&detail)
	if res.RowsAffected < 1 {
		return detail, fmt.Errorf("not record")
	}
	return detail, nil
}

// Update implements domain.CourseDetailRepository
func (dr *detailRepo) Update(id int, detail model.CourseDetail) (res model.CourseDetail, err error) {
	data := dr.DB.Where("course_id = ?", id).UpdateColumns(&detail)
	if data.RowsAffected < 1 {
		return detail, fmt.Errorf("error update data")
	}
	return detail, nil
}

func NewCourseDetailRepository(db *gorm.DB) domain.CourseDetailRepository {
	return &detailRepo{
		DB: db,
	}
}
