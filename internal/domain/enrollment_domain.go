package domain

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type EnrollmenRepository interface {
	Create(enr model.Enrollment) error
	GetByCourse(course_id int) (enr []model.Enrollment, err error)
	GetByUser(user_id int) (enr []model.Enrollment, err error)
	Update(user_id int, course_id int, enr model.Enrollment) error
	UpdateStatus(user_id int, course_id int, enr model.Enrollment) error
}

type EnrollmentService interface {
	Join(enr model.Enrollment, code string) error
	GetAllByCourseID(courseid int) (enr []model.Enrollment, err error)
	GetAllByUser(id int) ([]model.Enrollment, error)
	Update(user_id int, course_id int, enr model.Enrollment) error
}
