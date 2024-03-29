package restecho

import (
	"log"
	"net/http"
	"strconv"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type UserCourseController struct {
	Service    domain.CourseService
	EService   domain.EnrollmentService
	CatService domain.CourseCategoryService
	RService   domain.RatingService
}

func (cc *UserCourseController) JoinCourse(c echo.Context) error {
	course_id, _ := strconv.Atoi(c.Param("course_id"))

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	user_id := int(claim["id"].(float64))

	code := JoinRequest{}
	c.Bind(&code)

	validate := validator.New()
	err := validate.Struct(code)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "data harus sesuai tidak boleh kosong",
			"status":   http.StatusBadRequest,
		})
	}

	enr := model.Enrollment{
		UserID:   user_id,
		CourseID: course_id,
	}

	err = cc.EService.Join(enr, code.Code)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "gagal join kelas",
			"reason":   err,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
	})
}

func (cc *UserCourseController) GetAll(c echo.Context) error {
	courses := []UserDashboardCourse{}

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	user_id := int(claim["id"].(float64))

	res, err := cc.EService.GetAllByUser(user_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error",
			"reason":   err,
		})
	}

	for i := range res {
		course, _ := cc.Service.GetOneCourse(res[i].CourseID)

		totalMember := len(course.Student)
		totalMaterial := len(course.Material)

		progress := (float64(res[i].Progress) / float64(totalMaterial)) * 100

		udc := UserDashboardCourse{
			TotalMember:   totalMember,
			TotalMaterial: totalMaterial,
			Progress:      int(progress),
		}

		courses = append(courses, getUserCourse(course, udc))
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"messages": "berhasil",
		"data":     courses,
	})
}

func (cc *UserCourseController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("course_id"))

	res, err := cc.Service.GetOneCourse(id)

	catId := res.CourseDetail.CategoryID

	catData := cc.CatService.GetOneCategory(catId)

	courseResponse := CourseResponseDashUser{}
	courseResponse.Category = catData.Category

	count := 0

	for i := range res.Material {
		courseResponse.Material = append(courseResponse.Material, getMaterial(res.Material[i]))
		if res.Material[i].Video != "" {
			count++
		}
	}

	rate, _ := cc.RService.GetByCourse(id)
	totalRate := len(rate)
	jumlahRate := 0.0
	for j := range rate {
		jumlahRate += rate[j].Rating
	}

	avgRate := jumlahRate / float64(totalRate)

	courseResponse.Rating = avgRate

	for j := range res.Assignment {
		courseResponse.Assignment = append(courseResponse.Assignment, getAssignment(res.Assignment[j]))
	}

	log.Println(courseResponse.Assignment)
	log.Println(1)

	courseResponse.TotalVideo = count
	courseResponse.TotalMember = len(res.Student)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or deleted",
		})
	}
	log.Println(2)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"data":     getCourseDashUser(res, courseResponse),
	})
}

func (cc *UserCourseController) Dashboard(c echo.Context) error {
	courses := []UserDashboardCourse{}
	asResp := []AssignmentResponseDash{}
	data := UserDashboardResponse{}

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	user_id := int(claim["id"].(float64))

	res, err := cc.EService.GetAllByUser(user_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error",
			"reason":   err,
		})
	}

	for i := range res {
		course, _ := cc.Service.GetOneCourse(res[i].CourseID)

		totalMember := len(course.Student)
		totalMaterial := len(course.Material)

		progress := (float64(res[i].Progress) / float64(totalMaterial)) * 100

		udc := UserDashboardCourse{
			TotalMember:   totalMember,
			TotalMaterial: totalMaterial,
			Progress:      int(progress),
		}

		courses = append(courses, getUserCourse(course, udc))

		catId := course.CourseDetail.CategoryID

		catData := cc.CatService.GetOneCategory(catId)

		for j := range course.Assignment {
			asData := AssignmentResponseDash{
				AssignmentMentorID: course.Assignment[j].AssignmentMentorID,
				Title:              course.Assignment[j].Title,
				CourseName:         course.CourseName,
				CoursCategory:      catData.Category,
			}
			asResp = append(asResp, asData)
		}
	}

	data = getUserDash(courses, asResp)

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"messages": "berhasil",
		"data":     data,
	})
}

func (cc *UserCourseController) Update(c echo.Context) error {
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	user_id := int(claim["id"].(float64))

	course_id, _ := strconv.Atoi(c.Param("course_id"))
	enr := model.Enrollment{}

	err := cc.EService.Update(user_id, course_id, enr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error",
			"reason":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "berhasil",
	})
}

func (cc *UserCourseController) GetHistory(c echo.Context) error {
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	user_id := int(claim["id"].(float64))

	res, err := cc.EService.GetAllByUser(user_id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "error",
			"reason":   err,
		})
	}

	courses := []model.Course{}
	data := HistoryResp{}

	for i := range res {
		if res[i].Status == "complete" {
			course, _ := cc.Service.GetOneCourse(res[i].CourseID)
			courses = append(courses, course)
		}
	}
	data.Course = courses
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "berhasil",
		"data":     data,
	})
}
