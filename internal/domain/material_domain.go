package domain

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type MaterialRepository interface {
	Create(material model.Material) error
	Update(id int, material model.Material) error
	Delete(id int) error
	GetAll(courseid int) []model.Material
	GetByID(id int) (material model.Material, err error)
}

type MaterialService interface {
	Store(material model.Material) (int, error)
	Edit(id int, material model.Material) error
	Delete(id int) error
	GetAllByCourseID(courseid int) []model.Material
	GetOneMaterial(id int) (model.Material, error)
}
