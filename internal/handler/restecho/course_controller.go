package restecho

import (
	"net/http"
	"strconv"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CourseController struct {
	service domain.CourseService
}

func (cc *CourseController) CreateCourse(c echo.Context) error {
	course := model.Course{}

	c.Bind(&course)
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	course.MentorID = int(claim["id"].(float64))

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

func (cc *CourseController) EditCourse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	course := model.Course{}
	c.Bind(&course)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	err := cc.service.Edit(id, int(claim["id"].(float64)), course)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "edited",
		"id":        id,
		"mentor id": claim["id"],
	})
}

func (cc *CourseController) DeleteCourse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	err := cc.service.Delete(id, int(claim["id"].(float64)))

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "cannot delete course | no id | unauthorized",
			"status":   http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success delete course",
	})
}

func (cc *CourseController) GetCourses(c echo.Context) error {
	courses := cc.service.GetAllCourses()
	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    courses,
	}, "  ")
}

func (cc *CourseController) GetCourse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := cc.service.GetOneCourse(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or deleted",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    res,
	})

}
