package restecho

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type Mentor struct {
	Name        string
	Job         string
	Description string
}

type MaterialResponse struct {
	Title       string
	Description string
	Video       string
	PPT         string
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
	Material    []MaterialResponse
}

type CreateResponse struct {
	CourseName  string
	Description string
	Avatar      string
}

func createResponse(model model.Course) CreateResponse {
	return CreateResponse{
		CourseName:  model.CourseName,
		Description: model.CourseDetail.Description,
		Avatar:      model.CourseDetail.Avatar,
	}
}

func getCourses(model model.Course) CoursesResponse {
	return CoursesResponse{
		CourseID:     model.CourseID,
		CourseName:   model.CourseName,
		CourseMentor: model.User.Name,
	}
}

func getCourse(model model.Course, material []MaterialResponse) CourseResponse {
	return CourseResponse{
		ID:          model.CourseID,
		Title:       model.CourseName,
		Description: model.CourseName,
		Mentor: Mentor{
			Name:        model.User.Name,
			Job:         model.User.Name,
			Description: model.User.Name,
		},
		Material: material,
	}
}

func getMaterial(model model.Material) MaterialResponse {
	return MaterialResponse{
		Title:       model.MaterialName,
		Description: model.MaterialName,
		Video:       model.Video,
		PPT:         model.PPT,
	}
}
