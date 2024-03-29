package restecho

import (
	"net/http"
	"strconv"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CourseCategoryController struct {
	Service domain.CourseCategoryService
}

func (cc *CourseCategoryController) CreateCourseCategory(c echo.Context) error {
	//cek role
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	if int(claim["role"].(float64)) != 1 { //role 1 = admin
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
			"status":   http.StatusUnauthorized,
		})
	}

	course := model.CourseCategory{}

	c.Bind(&course)
	// bearer := c.Get("user").(*jwt.Token)
	// claim := bearer.Claims.(jwt.MapClaims)
	// course.MentorID = int(claim["id"].(float64))

	rescode, err := cc.Service.Store(course)

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

func (cc *CourseCategoryController) GetCourseCategory(c echo.Context) error {
	id := c.Param("course_category_id")
	intID, _ := strconv.Atoi(id)
	courses := cc.Service.GetOneCategory(intID)
	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    courses,
	}, "  ")
}
func (cc *CourseCategoryController) GetCoursesCategory(c echo.Context) error {
	courses := cc.Service.GetAllCategory()
	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    courses,
	}, "  ")
}
