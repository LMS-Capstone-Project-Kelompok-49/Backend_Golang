package restecho

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/bucket"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type CourseController struct {
	Service domain.CourseService
}

func (cc *CourseController) CreateCourse(c echo.Context) error {

	temp := CourseRequest{}

	c.Bind(&temp)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	mentorID := int(claim["id"].(float64))

	validate := validator.New()
	err := validate.Struct(temp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "data harus sesuai tidak boleh kosong",
			"status":   http.StatusBadRequest,
		})
	}

	currentTime := time.Now().Format("2006#01#02#15#04#05")
	code := strings.ReplaceAll(currentTime, "#", "")

	avatar, err := c.FormFile("avatar")
	if err != nil {
		temp.Avatar = ""
	} else {
		url, err := bucket.Open(avatar, "avatar", fmt.Sprintf("ava_cm%d_%s", mentorID, code))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "error upload file",
			})
		}

		temp.Avatar = url
	}

	course := model.Course{}
	course = toModelCourse(temp)
	course.CourseDetail = toModelDetail(temp)

	course.MentorID = mentorID

	rescode, err := cc.Service.Store(course)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
			"status":   rescode,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":   rescode,
		"messages": "success",
		"data":     createResponse(course),
	})
}

func (cc *CourseController) EditCourse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("course_id"))
	course := model.Course{}
	c.Bind(&course)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	cekMentor, err := cc.Service.GetOneCourse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or deleted",
		})
	}

	if int(claim["id"].(float64)) != cekMentor.User.UserID || int(claim["role"].(float64)) != 1 { //role 1 = admin
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
		})
	}

	err = cc.Service.Edit(id, course)
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
	id, _ := strconv.Atoi(c.Param("course_id"))

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	_, err := cc.Service.GetOneCourse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or deleted",
		})
	}

	if int(claim["role"].(float64)) != 1 { //role 1 = admin
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
		})
	}

	err = cc.Service.Delete(id)

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
	courses := cc.Service.GetAllCourses()
	data := []CoursesResponse{}
	for i := range courses {
		data = append(data, getCourses(courses[i]))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   http.StatusOK,
		"messages": "berhasil",
		"data":     data,
	})
}

func (cc *CourseController) GetCourse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("course_id"))

	res, err := cc.Service.GetOneCourse(id)

	material := []MaterialResponse{}

	for i := range res.Material {
		material = append(material, getMaterial(res.Material[i]))
	}

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or deleted",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"data":     getCourse(res, material),
	})
}
