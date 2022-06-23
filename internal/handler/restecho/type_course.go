package restecho

import (
	"net/http"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/labstack/echo/v4"
)

type TypeCategoryController struct {
	service domain.TypeCategoryService
}

func (cc *TypeCategoryController) GetTypeCourse(c echo.Context) error {
	courses := cc.service.GetAllType()
	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    courses,
	}, "  ")
}
