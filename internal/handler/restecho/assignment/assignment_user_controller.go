package restecho

import (
	"log"
	"net/http"
	"strconv"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AssignmentUserController struct {
	Service  domain.AssignmentUserService
	MService domain.AssignmentMentorService
}

func (ac *AssignmentUserController) CreateAssignment(c echo.Context) error {
	tugas := model.AssignmentUser{}
	amID, _ := strconv.Atoi(c.Param("am_id"))
	c.Bind(&tugas)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	user_id := int(claim["id"].(float64))

	tugas.UserID = user_id
	tugas.AssignmentMentorID = amID

	data, err := ac.Service.Store(tugas)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error post assignment",
			"status":  http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success post assignment",
		"status":  http.StatusOK,
		"data":    data,
	})
}

func (ac *AssignmentUserController) GetAssignmentByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	soal, err := ac.MService.GetOne(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "no data or deleted",
		})
	}

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	user_id := int(claim["id"].(float64))

	assignmentID := 0

	for i := range soal.AssignmentUser {
		if soal.AssignmentUser[i].UserID == user_id {
			assignmentID = soal.AssignmentUser[i].AssignmentUserID
		}
	}

	data, err := ac.Service.GetOne(assignmentID)
	if err != nil {
		log.Println("belum ngumpulin")
	}
	myAssignment := data.Document

	if myAssignment == "" {
		myAssignment = "kirim jawaban"
	}

	res := AssignmentResponse{
		Title:        soal.Title,
		Point:        soal.Point, //sementara
		Intruction:   soal.Intruction,
		Refference:   soal.Refference,
		MyAssignment: myAssignment,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success post assignment",
		"status":  http.StatusOK,
		"data":    getAssignment(res),
	})
}

func (ac *AssignmentUserController) GetAssignmentByUserID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("user_id"))

	res, err := ac.Service.GetAllByUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "no data or deleted",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success post assignment",
		"status":  http.StatusOK,
		"data":    res,
	})
}
