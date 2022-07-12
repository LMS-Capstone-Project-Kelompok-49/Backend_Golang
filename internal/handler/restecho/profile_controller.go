package restecho

import (
	"fmt"
	"net/http"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	service domain.ProfileService
}

func (pc *ProfileController) CreateProfile(c echo.Context) error {
	profile := model.UserProfile{}

	c.Bind(&profile)
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	profile.UserID = int(claim["id"].(float64))

	rescode, err := pc.service.Store(profile)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
			"status":   rescode,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success " + fmt.Sprintf("%d", rescode),
		"profile":  rescode,
	})
}

func (pc *ProfileController) EditProfile(c echo.Context) error {
	profile := model.UserProfile{}
	c.Bind(&profile)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	user_id := int(claim["id"].(float64))

	err := pc.service.Edit(user_id, profile)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (pc *ProfileController) GetProfile(c echo.Context) error {
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	user_id := int(claim["id"].(float64))
	res, err := pc.service.GetProfile(user_id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "profile empty, create first",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    res,
	})

}
