package restecho

import (
	"net/http"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/labstack/echo/v4"
)

type CourseCategoryController struct {
	service domain.CourseCategoryService
}

func (cc *CourseCategoryController) CreateCourseCategory(c echo.Context) error {
	course := model.CourseCategory{}

	c.Bind(&course)
	// bearer := c.Get("user").(*jwt.Token)
	// claim := bearer.Claims.(jwt.MapClaims)
	// course.MentorID = int(claim["id"].(float64))

	rescode, err := cc.service.Store(course)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
			"status":   rescode,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success " + string(rescode),
		"users":    course,
	})
}

func (cc *CourseCategoryController) GetCoursesCategory(c echo.Context) error {
	courses := cc.service.GetAllCategory()
	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    courses,
	}, "  ")
}
