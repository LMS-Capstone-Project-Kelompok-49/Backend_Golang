package restecho

import (
	"net/http"
	"strconv"

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

func (cc *TypeCategoryController) GetOneTypeCourse(c echo.Context) error {
	id := c.Param("type_course_id")
	intID, _ := strconv.Atoi(id)

	courses := cc.service.GetOneType(intID)
	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    courses,
	}, "  ")
}
