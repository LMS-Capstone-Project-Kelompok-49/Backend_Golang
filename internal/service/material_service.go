package service

import (
	"net/http"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gopkg.in/go-playground/validator.v9"
)

type materialService struct {
	repo domain.MaterialRepository
}

// Delete implements domain.MaterialService
func (ms *materialService) Delete(id int) error {
	return ms.repo.Delete(id)
}

// Edit implements domain.MaterialService
func (ms *materialService) Edit(id int, material model.Material) error {
	return ms.repo.Update(id, material)
}

// GetAllByCourseID implements domain.MaterialService
func (ms *materialService) GetAllByCourseID(courseid int) []model.Material {
	return ms.repo.GetAll(courseid)
}

// GetOneCourse implements domain.MaterialService
func (ms *materialService) GetOneMaterial(id int) (model.Material, error) {
	return ms.repo.GetByID(id)
}

// Store implements domain.MaterialService
func (ms *materialService) Store(material model.Material) (int, error) {
	validate := validator.New()
	err := validate.Struct(material)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, ms.repo.Create(material)
}

func NewMaterialService(repo domain.MaterialRepository) domain.MaterialService {
	return &materialService{
		repo: repo,
	}
}
