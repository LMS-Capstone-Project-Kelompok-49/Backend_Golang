package restecho

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type MentorResponse struct {
	Name        string `json:"name"`
	Job         string `json:"job"`
	Description string `json:"description"`
}

type MaterialResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Video       string `json:"video"`
	PPT         string `json:"ppt"`
}

type CoursesResponse struct {
	CourseID       int     `json:"course_id"`
	CourseName     string  `json:"course_name"`
	CourseMentor   string  `json:"course_mentor"`
	CourseCategory string  `json:"course_category"`
	Rating         float32 `json:"rating"`
	TotalVideo     int     `json:"total_video"`
	TotalMember    int     `json:"total_member"`
	Price          int     `json:"price"`
}

type CourseResponse struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Category    string             `json:"category"`
	Media       string             `json:"media"`
	Mentor      MentorResponse     `json:"mentor"`
	Material    []MaterialResponse `json:"material"`
	Benefit     string             `json:"benefit"`
	Rating      string             `json:"rating"`
	TotalVideo  int                `json:"total_video"`
	Price       int                `json:"price"`
}

type CreateResponse struct {
	CourseName  string `json:"course_name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

func createResponse(model model.Course) CreateResponse {
	return CreateResponse{
		CourseName:  model.CourseName,
		Description: model.CourseDetail.Description,
		Avatar:      model.CourseDetail.Avatar,
	}
}

func getCourses(model model.Course, cr CoursesResponse) CoursesResponse {
	return CoursesResponse{
		CourseID:       model.CourseID,
		CourseName:     model.CourseName,
		CourseMentor:   model.User.Profile.Fullname,
		CourseCategory: cr.CourseCategory,
		TotalVideo:     cr.TotalVideo,
	}
}

func getCourse(model model.Course, cr CourseResponse) CourseResponse {
	return CourseResponse{
		ID:          model.CourseID,
		Title:       model.CourseName,
		Description: model.CourseDetail.Description,
		Media:       model.CourseDetail.Media,
		Category:    cr.Category,
		Mentor: MentorResponse{
			Name:        model.User.Profile.Fullname,
			Job:         model.User.Profile.Job,
			Description: model.User.Profile.Description,
		},
		Material:   cr.Material,
		TotalVideo: cr.TotalVideo,
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
