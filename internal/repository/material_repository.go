package repository

import (
	"fmt"
	"log"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type materialRepoLayer struct {
	DB *gorm.DB
}

// Create implements domain.MaterialRepository
func (mr *materialRepoLayer) Create(material model.Material) error {
	re := mr.DB.Model(&model.Course{}).Association("CourseID")
	log.Println(re)
	res := mr.DB.Create(&material)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create material")
	}
	return nil
}

// Delete implements domain.MaterialRepository
func (mr *materialRepoLayer) Delete(id int) error {
	res := mr.DB.Delete(&model.Material{
		MaterialID: id,
	})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete material")
	}
	return nil
}

// GetAll implements domain.MaterialRepository
func (mr *materialRepoLayer) GetAll(courseid int) (materials []model.Material, err error) {
	res := mr.DB.Where("course_id = ?", courseid).Find(&materials)
	if res.RowsAffected < 1 {
		return materials, fmt.Errorf("error delete material")
	}
	return materials, nil
}

// GetByID implements domain.MaterialRepository
func (mr *materialRepoLayer) GetByID(id int) (material model.Material, err error) {
	err = mr.DB.Where("material_id = ?", id).First(&material).Error
	if err != nil {
		return material, fmt.Errorf("not found")
	}
	return material, nil
}

// Update implements domain.MaterialRepository
func (mr *materialRepoLayer) Update(id int, material model.Material) error {
	res := mr.DB.Where("material_id = ?", id).UpdateColumns(&material)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update data")
	}
	return nil
}

func NewMaterialRepository(db *gorm.DB) domain.MaterialRepository {
	return &materialRepoLayer{
		DB: db,
	}
}
