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

type AssignmentResponse struct {
	AssignmentMentorID int    `json:"assigment_id"`
	Title              string `json:"title"`
	Refference         string `json:"refference"`
	Point              int    `json:"point"`
}

type CoursesResponse struct {
	CourseID       int     `json:"course_id"`
	CourseName     string  `json:"course_name"`
	CourseMentor   string  `json:"course_mentor"`
	CourseCategory string  `json:"course_category"`
	Rating         float64 `json:"rating"`
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
	Rating      float64            `json:"rating"`
	TotalVideo  int                `json:"total_video"`
	TotalMember int                `json:"total_member"`
	Price       int                `json:"price"`
}

type CourseResponseDash struct {
	ID          int                  `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Category    string               `json:"category"`
	Media       string               `json:"media"`
	Mentor      MentorResponse       `json:"mentor"`
	Material    []MaterialResponse   `json:"material"`
	Assignment  []AssignmentResponse `json:"assignment"`
	Benefit     string               `json:"benefit"`
	Rating      float64              `json:"rating"`
	TotalVideo  int                  `json:"total_video"`
	TotalMember int                  `json:"total_member"`
	Price       int                  `json:"price"`
}

type CreateResponse struct {
	CourseName  string `json:"course_name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

type UserDashboardCourse struct {
	CourseID      int
	CourseName    string
	Mentor        string
	TotalMember   int
	TotalMaterial int
	Progress      int
}

func getUserCourse(model model.Course, udc UserDashboardCourse) UserDashboardCourse {
	return UserDashboardCourse{
		CourseID:      model.CourseID,
		CourseName:    model.CourseName,
		Mentor:        model.Mentor.Profile.Fullname,
		TotalMember:   udc.TotalMember,
		TotalMaterial: udc.TotalMaterial,
		Progress:      udc.Progress,
	}
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
		CourseMentor:   model.Mentor.Profile.Fullname,
		CourseCategory: cr.CourseCategory,
		TotalVideo:     cr.TotalVideo,
		Rating:         cr.Rating,
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
			Name:        model.Mentor.Profile.Fullname,
			Job:         model.Mentor.Profile.Job,
			Description: model.Mentor.Profile.Description,
		},
		Material:    cr.Material,
		TotalVideo:  cr.TotalVideo,
		TotalMember: cr.TotalMember,
	}
}

func getCourseDash(model model.Course, cr CourseResponseDash) CourseResponseDash {
	return CourseResponseDash{
		ID:          model.CourseID,
		Title:       model.CourseName,
		Description: model.CourseDetail.Description,
		Media:       model.CourseDetail.Media,
		Category:    cr.Category,
		Mentor: MentorResponse{
			Name:        model.Mentor.Profile.Fullname,
			Job:         model.Mentor.Profile.Job,
			Description: model.Mentor.Profile.Description,
		},
		Material:    cr.Material,
		Assignment:  cr.Assignment,
		TotalVideo:  cr.TotalVideo,
		TotalMember: cr.TotalMember,
	}
}

func getMaterial(model model.Material) MaterialResponse {
	return MaterialResponse{
		Title:       model.MaterialName,
		Description: model.Description,
		Video:       model.Video,
		PPT:         model.PPT,
	}
}

func getAssignment(model model.AssignmentMentor) AssignmentResponse {
	return AssignmentResponse{
		AssignmentMentorID: model.AssignmentMentorID,
		Title:              model.Title,
		Refference:         model.Refference,
		Point:              model.Point,
	}
}
