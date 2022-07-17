package restecho

import (
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
)

type CourseRequest struct {
	CourseName  string `json:"coursename" form:"coursename" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Avatar      string `json:"avatar" form:"avatar"`
}

func toModelCourse(cr CourseRequest) model.Course {
	return model.Course{
		CourseName: cr.CourseName,
	}
}

func toModelDetail(cr CourseRequest) model.CourseDetail {
	return model.CourseDetail{
		Description: cr.Description,
		Avatar:      cr.Avatar,
	}
}
