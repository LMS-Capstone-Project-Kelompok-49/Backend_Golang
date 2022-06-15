package restecho

import (
	"net/http"
	"strconv"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RoleController struct {
	service  domain.RoleService
	uservice domain.UserAdapterService
}

func (rc *RoleController) CreateRole(c echo.Context) error {
	role := model.Role{}

	c.Bind(&role)
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	usr, _ := rc.uservice.GetUserByID(int(claim["id"].(float64)))

	if usr.RoleID != 1 {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
		})
	}

	rescode, err := rc.service.Store(role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
			"status":   rescode,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success ",
		"users":    role,
	})
}

func (rc *RoleController) EditRole(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	role := model.Role{}
	c.Bind(&role)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	usr, _ := rc.uservice.GetUserByID(int(claim["id"].(float64)))

	if usr.RoleID != 1 {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
		})
	}

	err := rc.service.Edit(id, role)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
		"id":       id,
	})
}

func (rc *RoleController) DeleteRole(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	usr, _ := rc.uservice.GetUserByID(int(claim["id"].(float64)))

	if usr.RoleID != 1 {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
		})
	}

	err := rc.service.Delete(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "cannot delete role | no id | unauthorized",
			"status":   http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success delete role",
	})
}

func (rc *RoleController) GetRole(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := rc.service.GetOneRole(id)

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
