package restecho

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/bucket"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
)

type MaterialController struct {
	Service       domain.MaterialService
	CourseService domain.CourseService
}

func (mc *MaterialController) CreateMaterial(c echo.Context) error {
	material := model.Material{}
	courseid, _ := strconv.Atoi(c.Param("courseid"))
	c.Bind(&material)

	//cek mentor
	// mentor, err := mc.CourseService.GetOneCourse(courseid)
	// bearer := c.Get("user").(*jwt.Token)
	// claim := bearer.Claims.(jwt.MapClaims)

	// if err != nil {
	// 	return c.JSON(http.StatusNotFound, map[string]interface{}{
	// 		"messages": "no id or no change or unauthorization",
	// 	})
	// }

	// if mentor.MentorID != int(claim["id"].(float64)) {
	// 	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
	// 		"messages": "unauthorized",
	// 		"status":   http.StatusUnauthorized,
	// 	})
	// }

	currentTime := time.Now().Format("2006#01#02#15#04#05")
	code := strings.ReplaceAll(currentTime, "#", "")

	video, err := c.FormFile("video")
	if err != nil {
		material.Video = ""
	} else {
		url, err := bucket.Open(video, "video", fmt.Sprintf("vid_c%d_%s", courseid, code))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "error upload file",
			})
		}

		material.Video = url
	}
	ppt, err := c.FormFile("ppt")
	if err != nil {
		material.PPT = ""
	} else {
		url, err := bucket.Open(ppt, "ppt", fmt.Sprintf("ppt_c%d_%s", courseid, code))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "error upload file",
			})
		}

		material.PPT = url
	}

	material.CourseID = courseid

	rescode, err := mc.Service.Store(material)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error tambah material",
			"status":   rescode,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success " + fmt.Sprintf("%d", rescode),
		"material": material,
	})
}

func (mc *MaterialController) EditMaterial(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := mc.Service.GetOneMaterial(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": err.Error(),
			"aa":       err.Error(),
		})
	}

	//cek mentor
	courseid := data.CourseID

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	mentor, err := mc.CourseService.GetOneCourse(courseid)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	if mentor.User.UserID != int(claim["id"].(float64)) {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
			"status":   http.StatusUnauthorized,
		})
	}

	material := model.Material{}
	c.Bind(&material)

	currentTime := time.Now().Format("2006#01#02#15#04#05")
	code := strings.ReplaceAll(currentTime, "#", "")

	video, err := c.FormFile("video")
	if err != nil {
		material.Video = ""
	} else {
		url, err := bucket.Open(video, "video", fmt.Sprintf("vid_c%d_%s", courseid, code))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "error upload file",
			})
		}

		material.Video = url
	}
	ppt, err := c.FormFile("ppt")
	if err != nil {
		material.PPT = ""
	} else {
		url, err := bucket.Open(ppt, "ppt", fmt.Sprintf("ppt_c%d_%s", courseid, code))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "error upload file",
			})
		}

		material.PPT = url
	}

	// bearer := c.Get("user").(*jwt.Token)
	// claim := bearer.Claims.(jwt.MapClaims)

	err1 := mc.Service.Edit(id, material)
	if err1 != nil {
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

	data, err := mc.Service.GetOneMaterial(id)
	if err != nil {
		return err
	}

	//cek mentor
	courseid := data.CourseID
	mentor, err := mc.CourseService.GetOneCourse(courseid)
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	if mentor.User.UserID != int(claim["id"].(float64)) {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
			"status":   http.StatusUnauthorized,
		})
	}

	if data.Video != "" {
		_, fname := filepath.Split(data.Video)
		err := bucket.RemoveFile(fname, "video")
		if err != nil {
			return err
		}
	}

	if data.PPT != "" {
		_, fname := filepath.Split(data.PPT)
		err := bucket.RemoveFile(fname, "ppt")
		if err != nil {
			return err
		}
	}

	// bearer := c.Get("user").(*jwt.Token)
	// claim := bearer.Claims.(jwt.MapClaims)

	err = mc.Service.Delete(id)

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
	materials, err := mc.Service.GetAllByCourseID(courseid)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or deleted",
		})
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages":  "success",
		"materials": materials,
	}, "  ")
}

func (mc *MaterialController) GetMaterial(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := mc.Service.GetOneMaterial(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or deleted",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"material": res,
	})
}
