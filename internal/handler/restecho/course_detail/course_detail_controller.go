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

type CourseDetailController struct {
	Service domain.CourseDetailService
}

func (cd *CourseDetailController) EditDetail(c echo.Context) error {
	detail := model.CourseDetail{}
	id, _ := strconv.Atoi(c.Param("course_id"))

	c.Bind(&detail)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	mentorID := int(claim["id"].(float64))

	currentTime := time.Now().Format("2006#01#02#15#04#05")
	code := strings.ReplaceAll(currentTime, "#", "")

	avatar, err := c.FormFile("avatar")
	if err != nil {
		detail.Avatar = ""
	} else {
		url, err := bucket.Open(avatar, "avatar", fmt.Sprintf("ava_cm%d_%s", mentorID, code))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "error upload file",
			})
		}

		detail.Avatar = url
	}

	media, err := c.FormFile("media")
	if err != nil {
		detail.Media = ""
	} else {
		url, err := bucket.Open(media, "media", fmt.Sprintf("om_cm%d_%s", mentorID, code))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "error upload file",
			})
		}

		detail.Media = url
	}

	res, err := cd.Service.Edit(id, detail)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "bad request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "berhasil update detail",
		"data":     res,
	})

}

func (cd *CourseDetailController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("course_id"))

	res, err := cd.Service.GetOne(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no data or id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "berhasil",
		"data":     res,
	})
}
