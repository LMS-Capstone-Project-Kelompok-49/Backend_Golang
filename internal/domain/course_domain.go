package domain

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type CourseRepository interface {
	Create(course model.Course) error
	Update(id int, course model.Course) error
	Delete(id int) error
	GetAll() []model.Course
	GetByID(id int) (course model.Course, err error)
}

type CourseService interface {
	Store(course model.Course) (int, error)
	Edit(id int, idToken int, course model.Course) error
	Delete(id int, idToken int) error
	GetAllCourses() []model.Course
	GetOneCourse(id int) (model.Course, error)
}
