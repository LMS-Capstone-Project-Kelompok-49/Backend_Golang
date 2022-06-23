package restecho

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/bucket"
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
	c.Bind(&material)
	video, err := c.FormFile("video")
	if err != nil {
		return fmt.Errorf("error")
	}
	if video != nil {
		vidSrc, err := video.Open()
		if err != nil {
			return fmt.Errorf("error")
		}
		defer vidSrc.Close()

		vidPath := filepath.Join("temp", video.Filename)

		vidDst, err := os.Create(vidPath)
		if err != nil {
			return fmt.Errorf("error")
		}
		defer vidDst.Close()

		if _, err := io.Copy(vidDst, vidSrc); err != nil {
			return fmt.Errorf("error")
		}

		if videoUrl, err := bucket.UploadFile(video.Filename, vidPath, "video"); err != nil {
			return fmt.Errorf("error")
		} else {
			material.Video = videoUrl
			err := os.Remove(vidPath)
			if err != nil {
				return fmt.Errorf("error")
			}
		}
	}
	ppt, err := c.FormFile("ppt")
	if err != nil {
		return fmt.Errorf("error")
	}
	if ppt != nil {
		pptSrc, err := ppt.Open()
		if err != nil {
			return fmt.Errorf("error")
		}
		defer pptSrc.Close()

		pptPath := filepath.Join("temp", ppt.Filename)

		pptDst, err := os.Create(pptPath)
		if err != nil {
			return fmt.Errorf("error")
		}
		defer pptDst.Close()

		if _, err := io.Copy(pptDst, pptSrc); err != nil {
			return fmt.Errorf("error")
		}

		if pptUrl, err := bucket.UploadFile(ppt.Filename, pptPath, "ppt"); err != nil {
			return fmt.Errorf("error")
		} else {
			material.PPT = pptUrl
			err := os.Remove(pptPath)
			if err != nil {
				return fmt.Errorf("error")
			}
		}
	}

	material.CourseID = courseid

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
