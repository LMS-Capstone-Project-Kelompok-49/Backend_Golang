package restecho

import (
	"net/http"
	"strconv"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type EchoController struct {
	Svc domain.UserAdapterService
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

	err, statusCode := ce.Svc.CreateUserService(user)

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
	id := c.Param("user_id")
	intID, _ := strconv.Atoi(id)

	user := model.User{}
	c.Bind(&user)

	//cek id || role
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	// idnya := int(claim["role"].(float64))
	// fmt.Println(idnya)
	// fmt.Printf("Roles: %d -- %d/n", idnya, int(claim["id"].(float64)))

	// if int(claim["id"].(float64)) != intID || int(claim["role"].(float64)) != 1 { //role 1 = admin
	if int(claim["role"].(float64)) != 1 { //role 1 = admin
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
			"status":   http.StatusUnauthorized,
		})
	}

	err := ce.Svc.UpdateUserService(intID, user)
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
	id := c.Param("user_id")
	intID, _ := strconv.Atoi(id)

	//cek id || role
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	if int(claim["role"].(float64)) != 1 { //role 1 = admin
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
			"status":   http.StatusUnauthorized,
		})
	}

	err := ce.Svc.DeleteByID(intID)
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
	id := c.Param("user_id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.Svc.GetUserByID(intID)
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

func (ce *EchoController) GetUsersController(c echo.Context) error {
	users := ce.Svc.GetAllUsersService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    users,
	}, "  ")
}

func (ce *EchoController) LoginUserController(c echo.Context) error {
	userLogin := make(map[string]interface{})

	c.Bind(&userLogin)

	token, statusCode := ce.Svc.LoginUser(userLogin["email"].(string), userLogin["password"].(string))
	switch statusCode {
	case http.StatusUnauthorized:
		return c.JSONPretty(http.StatusBadRequest, map[string]interface{}{
			"status":   http.StatusBadRequest,
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
