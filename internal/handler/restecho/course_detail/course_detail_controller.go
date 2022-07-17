package restecho

import (
	"net/http"
	"strconv"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/labstack/echo/v4"
)

type CourseDetailController struct {
	Service domain.CourseDetailService
}

func (cd *CourseDetailController) EditDetail(c echo.Context) error {
	detail := model.CourseDetail{}
	id, _ := strconv.Atoi(c.Param("course_id"))
	c.Bind(&detail)

	err := cd.Service.Edit(id, detail)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "bad request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "berhasil update detail",
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
