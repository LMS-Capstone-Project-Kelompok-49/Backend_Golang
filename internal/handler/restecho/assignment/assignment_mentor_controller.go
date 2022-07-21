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
)

type AssignmentMentorController struct {
	Service domain.AssignmentMentorService
}

func (ac *AssignmentMentorController) CreateAssignment(c echo.Context) error {
	tugas := model.AssignmentMentor{}
	course_id, _ := strconv.Atoi(c.Param("course_id"))
	c.Bind(&tugas)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	mentorID := int(claim["id"].(float64))

	currentTime := time.Now().Format("2006#01#02#15#04#05")
	code := strings.ReplaceAll(currentTime, "#", "")

	reff, err := c.FormFile("refference")
	if err != nil {
		tugas.Refference = ""
	} else {
		url, err := bucket.Open(reff, "assignmentmentor", fmt.Sprintf("reff_cm%d_%s", mentorID, code))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "error upload file",
			})
		}

		tugas.Refference = url
	}

	tugas.CourseID = course_id

	err = ac.Service.Store(tugas)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error menambahkan tugas",
			"reason":   err,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "berhasil",
	})

}

func (ac *AssignmentMentorController) EditAssignment(c echo.Context) error {
	tugas := model.AssignmentMentor{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&tugas)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	mentorID := int(claim["id"].(float64))

	currentTime := time.Now().Format("2006#01#02#15#04#05")
	code := strings.ReplaceAll(currentTime, "#", "")

	reff, err := c.FormFile("refference")
	if err != nil {
		tugas.Refference = ""
	} else {
		url, err := bucket.Open(reff, "assignmentmentor", fmt.Sprintf("reff_cm%d_%s", mentorID, code))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "error upload file",
			})
		}

		tugas.Refference = url
	}

	err = ac.Service.Edit(id, tugas)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error edit tugas",
			"reason":   err,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "berhasil",
	})

}

func (ac *AssignmentMentorController) DeleteAssignment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := ac.Service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error menghapus tugas",
			"reason":   err,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "berhasil",
	})

}

func (ac *AssignmentMentorController) GetAssignment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tugas, err := ac.Service.GetOne(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "no id or deleted",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "berhasil",
		"data":     tugas,
	})

}
