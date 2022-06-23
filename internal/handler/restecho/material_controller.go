package restecho

import (
	"net/http"
	"strconv"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/labstack/echo/v4"
)

type MaterialController struct {
	service domain.MaterialService
}

func (mc *MaterialController) CreateMaterial(c echo.Context) error {
	material := model.Material{}
	courseid, _ := strconv.Atoi(c.Param("courseid"))
	// file, err := c.FormFile("video")
	// if err != nil {
	// 	return fmt.Errorf("error")
	// }

	// // dir, err := os.Getwd()
	// // if err != nil {
	// // 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// // 		"messages": "err.Error()",
	// // 	})
	// // }

	// src, err := file.Open()
	// if err != nil {
	// 	return fmt.Errorf("error")
	// }
	// defer src.Close()

	// dst, err := os.Create(file.Filename)
	// if err != nil {
	// 	return err
	// }
	// defer dst.Close()

	// if _, err := io.Copy(dst, src); err != nil {
	// 	return err
	// }

	c.Bind(&material)
	// bearer := c.Get("user").(*jwt.Token)
	// claim := bearer.Claims.(jwt.MapClaims)
	// course.MentorID = int(claim["id"].(float64))
	material.CourseID = courseid

	//upload file
	// fmt.Println(file.Header)
	// res, err := gdrive.GoUp(file.Filename)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"messages": err.Error(),
	// 	})
	// }
	// material.Video = res.Id

	rescode, err := mc.service.Store(material)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
			"status":   rescode,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success " + string(rescode),
		"material": material,
	})
}

func (mc *MaterialController) EditMaterial(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	material := model.Material{}
	c.Bind(&material)

	// bearer := c.Get("user").(*jwt.Token)
	// claim := bearer.Claims.(jwt.MapClaims)

	err := mc.service.Edit(id, material)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (mc *MaterialController) DeleteMaterial(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// bearer := c.Get("user").(*jwt.Token)
	// claim := bearer.Claims.(jwt.MapClaims)

	err := mc.service.Delete(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "cannot delete material | no id | unauthorized",
			"status":   http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success delete material",
	})
}

func (mc *MaterialController) GetMaterials(c echo.Context) error {
	courseid, _ := strconv.Atoi(c.Param("courseid"))
	materials := mc.service.GetAllByCourseID(courseid)
	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages":  "success",
		"materials": materials,
	}, "  ")
}

func (mc *MaterialController) GetMaterial(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := mc.service.GetOneMaterial(id)

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
