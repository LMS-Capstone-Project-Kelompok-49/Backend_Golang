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
	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
)

type MaterialController struct {
	service       domain.MaterialService
	courseService domain.CourseService
}

func (mc *MaterialController) CreateMaterial(c echo.Context) error {
	material := model.Material{}
	courseid, _ := strconv.Atoi(c.Param("courseid"))
	c.Bind(&material)

	//cek mentor
	mentor, err := mc.courseService.GetOneCourse(courseid)
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	if mentor.MentorID != int(claim["id"].(float64)) {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
			"status":   http.StatusUnauthorized,
		})
	}

	video, err := c.FormFile("video")
	if err != nil {
		material.Video = ""
	} else {
		vidSrc, err := video.Open()
		if err != nil {
			return fmt.Errorf("error")
		}

		vidPath := filepath.Join("temp", video.Filename)

		vidDst, err := os.Create(vidPath)
		if err != nil {
			return fmt.Errorf("error")
		}

		if _, err := io.Copy(vidDst, vidSrc); err != nil {
			return fmt.Errorf("error")
		}

		if videoUrl, err := bucket.UploadFile(video.Filename, vidPath, "video"); err != nil {
			return fmt.Errorf("error")
		} else {
			vidSrc.Close()
			vidDst.Close()

			material.Video = videoUrl

			err := os.Remove(vidPath)
			if err != nil {
				return err
			}
		}
	}
	ppt, err := c.FormFile("ppt")
	if err != nil {
		material.PPT = ""
	} else {
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
			pptSrc.Close()
			pptDst.Close()

			material.PPT = pptUrl

			err := os.Remove(pptPath)
			if err != nil {
				return err
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
		"messages": "success " + fmt.Sprintf("%d", rescode),
		"material": material,
	})
}

func (mc *MaterialController) EditMaterial(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := mc.service.GetOneMaterial(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": err.Error(),
			"aa":       err.Error(),
		})
	}

	//cek mentor
	courseid := data.CourseID
	mentor, err := mc.courseService.GetOneCourse(courseid)
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	if mentor.MentorID != int(claim["id"].(float64)) {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "unauthorized",
			"status":   http.StatusUnauthorized,
		})
	}

	material := model.Material{}
	c.Bind(&material)

	video, err := c.FormFile("video")
	if err != nil {
		material.Video = ""
	} else {
		vidSrc, err := video.Open()
		if err != nil {
			return fmt.Errorf("error")
		}

		vidPath := filepath.Join("temp", video.Filename)

		vidDst, err := os.Create(vidPath)
		if err != nil {
			return fmt.Errorf("error")
		}

		if _, err := io.Copy(vidDst, vidSrc); err != nil {
			return fmt.Errorf("error")
		}

		if videoUrl, err := bucket.UploadFile(video.Filename, vidPath, "video"); err != nil {
			return fmt.Errorf("error")
		} else {
			vidSrc.Close()
			vidDst.Close()

			_, fname := filepath.Split(data.Video)

			err := bucket.RemoveFile(fname, "video")
			if err != nil {
				return err
			}

			material.Video = videoUrl

			err = os.Remove(vidPath)
			if err != nil {
				return err
			}
		}
	}
	ppt, err := c.FormFile("ppt")
	if err != nil {
		material.PPT = ""
	} else {
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
			pptSrc.Close()
			pptDst.Close()

			_, fname := filepath.Split(data.PPT)

			err := bucket.RemoveFile(fname, "ppt")
			if err != nil {
				return err
			}

			material.PPT = pptUrl

			err = os.Remove(pptPath)
			if err != nil {
				return err
			}
		}
	}

	// bearer := c.Get("user").(*jwt.Token)
	// claim := bearer.Claims.(jwt.MapClaims)

	err1 := mc.service.Edit(id, material)
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

	data, err := mc.service.GetOneMaterial(id)
	if err != nil {
		return err
	}

	//cek mentor
	courseid := data.CourseID
	mentor, err := mc.courseService.GetOneCourse(courseid)
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	if mentor.MentorID != int(claim["id"].(float64)) {
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

	err = mc.service.Delete(id)

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
