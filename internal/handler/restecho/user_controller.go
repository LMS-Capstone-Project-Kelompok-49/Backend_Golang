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

type EchoController struct {
	svc domain.UserAdapterService
}

// Get one godoc
// @Summary Create content
// @description create content with data
// @Security BearerAuth
// @tags content
// @Accept json
// @Produce json
// @Success 200 {object} content.Content
// @Router /app/u [get]
func (ce *EchoController) CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	user.RoleID = 2

	err, statusCode := ce.svc.CreateUserService(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
			"status":   statusCode,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success " + string(statusCode),
		"users":    user,
	})
}

func (ce *EchoController) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	user := model.User{}
	c.Bind(&user)

	//cek id || role
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	if int(claim["id"].(float64)) != intID || int(claim["role"].(float64)) != 1 { //role 1 = admin
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
			"status":   http.StatusUnauthorized,
		})
	}

	err := ce.svc.UpdateUserService(intID, user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":   "edited",
		"id":         intID,
		"expeted id": claim["id"],
	})
}

func (ce *EchoController) DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	//cek id || role
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	if int(claim["id"].(float64)) != intID || int(claim["role"].(float64)) != 1 { //role 1 = admin
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
			"status":   http.StatusUnauthorized,
		})
	}

	err := ce.svc.DeleteByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoController) GetUserController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.svc.GetUserByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    res,
	})
}

// Create godoc
// @Summary Create content
// @description create content with data
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @tags content
// @Accept json
// @Produce json
// @Success 200 {object} content.Content
// @Router /app [post]

func (ce *EchoController) GetUsersController(c echo.Context) error {
	users := ce.svc.GetAllUsersService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    users,
	}, "  ")
}

func (ce *EchoController) LoginUserController(c echo.Context) error {
	userLogin := make(map[string]interface{})

	c.Bind(&userLogin)

	token, statusCode := ce.svc.LoginUser(userLogin["email"].(string), userLogin["password"].(string))
	switch statusCode {
	case http.StatusUnauthorized:
		return c.JSONPretty(http.StatusUnauthorized, map[string]interface{}{
			"messages": "email atau password salah",
		}, "  ")

	case http.StatusInternalServerError:
		return c.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal",
		}, "  ")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"token":    token,
	}, "  ")
}
