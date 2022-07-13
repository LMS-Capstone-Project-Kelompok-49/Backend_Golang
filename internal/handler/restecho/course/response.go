package restecho

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type Mentor struct {
	Name        string
	Job         string
	Description string
}

type CoursesResponse struct {
	CourseID     int
	CourseName   string
	CourseTypeID string
	CourseMentor string
}

type CourseResponse struct {
	ID          int
	Title       string
	Description string
	Mentor      Mentor
}

func getCourses(model model.Course) CoursesResponse {
	return CoursesResponse{
		CourseID:     model.CourseID,
		CourseName:   model.CourseName,
		CourseTypeID: model.CourseType,
		CourseMentor: model.User.Name,
	}
}

func getCourse(model model.Course) CourseResponse {
	return CourseResponse{
		ID:          model.CourseID,
		Title:       model.CourseName,
		Description: model.CourseName,
		Mentor: Mentor{
			Name:        model.User.Name,
			Job:         model.User.Name,
			Description: model.User.Name,
		},
	}
}
