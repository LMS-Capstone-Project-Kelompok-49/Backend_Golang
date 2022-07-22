package service

import (
	"fmt"
	"log"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type enrollementService struct {
	repo  domain.CourseRepository
	erepo domain.EnrollmenRepository
}

// Update implements domain.EnrollmentService
func (es *enrollementService) Update(user_id int, course_id int, enr model.Enrollment) error {
	course, err := es.repo.GetByID(course_id)
	if err != nil {
		return fmt.Errorf("no course")
	}

	data, err := es.erepo.GetByUser(user_id)
	if err != nil {
		return fmt.Errorf("no enrollment")
	}

	totalMaterial := len(course.Material)
	log.Println(totalMaterial)
	totalProgress := 0

	for i := range data {
		if data[i].CourseID == course_id && data[i].UserID == user_id {
			totalProgress = data[i].Progress
		}
	}
	log.Println(totalProgress)

	selisih := totalMaterial - totalProgress
	log.Println(selisih)

	if selisih == 0 {
		return fmt.Errorf("progress sudah selesai")
	} else if selisih == 1 {
		return es.erepo.UpdateStatus(user_id, course_id, enr)
	} else {
		return es.erepo.Update(user_id, course_id, enr)
	}

}

// GetAllByCourseID implements domain.EnrollmentService
func (*enrollementService) GetAllByCourseID(courseid int) (enr []model.Enrollment, err error) {
	panic("unimplemented")
}

// GetAllByUser implements domain.EnrollmentService
func (es *enrollementService) GetAllByUser(id int) ([]model.Enrollment, error) {
	return es.erepo.GetByUser(id)
}

// Join implements domain.EnrollmentService
func (es *enrollementService) Join(enr model.Enrollment, code string) error {
	course, err := es.repo.GetByID(enr.CourseID)
	if err != nil {
		return fmt.Errorf("no course")
	}

	if code != course.Code {
		return fmt.Errorf("kode tidak berlaku")
	}

	return es.erepo.Create(enr)
}

func NewEnrollmentService(erepo domain.EnrollmenRepository, repo domain.CourseRepository) domain.EnrollmentService {
	return &enrollementService{
		erepo: erepo,
		repo:  repo,
	}
}
