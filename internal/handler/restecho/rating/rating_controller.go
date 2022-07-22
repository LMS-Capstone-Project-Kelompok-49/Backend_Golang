package restecho

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RatingController struct {
	Service       domain.RatingService
	CourseService domain.CourseService
}

// Store implements domain.RatingService
func (rc *RatingController) Create(c echo.Context) error {
	rating := model.Rating{}
	course_id, _ := strconv.Atoi(c.Param("course_id"))

	_, err := rc.CourseService.GetOneCourse(course_id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": fmt.Sprintf("no course with id %d", course_id),
		})
	}

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	usrID := int(claim["id"].(float64))

	c.Bind(&rating)
	rating.UserID = usrID
	rating.CourseID = course_id

	err = rc.Service.Store(rating)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "gagal tambah data",
			"reason":  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil tambah data",
	})
}

// Edit implements domain.RatingService
func (rc *RatingController) Update(c echo.Context) error {
	rating := model.Rating{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&rating)

	res, err := rc.Service.Edit(id, rating)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "gagal ubah data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil ubah data",
		"data":    res,
	})
}

// Remove implements domain.RatingService
func (rc *RatingController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := rc.Service.Remove(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "gagal hapus data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil hapus data",
	})
}

// GetByCourse implements domain.RatingService
func (rc *RatingController) GetByCourse(c echo.Context) error {
	course_id, _ := strconv.Atoi(c.Param("course_id"))
	res, err := rc.Service.GetByCourse(course_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "gagal ambil data | tidak ada id | terhapus",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil ambil data",
		"data":    res,
	})
}

// GetOne implements domain.RatingService
func (rc *RatingController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := rc.Service.GetOne(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "gagal ambil data | tidak ada id | terhapus",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil ambil data",
		"data":    res,
	})
}
