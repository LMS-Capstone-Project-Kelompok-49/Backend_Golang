package restecho

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/config"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/database"
	m "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/middleware"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/repository"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/service"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepository(db)

	svc := service.NewServiceUser(repo, conf)

	cont := EchoController{
		svc: svc,
	}
	auth := e.Group("/auth",
		middleware.Logger(),
		middleware.CORS(),
	)
	apiUser := e.Group("/app",
		middleware.Logger(),
		middleware.CORS(),
		m.APIKEYMiddleware,
	)
	// Auth Handler
	auth.POST("/login", cont.LoginUserController)
	auth.POST("/register", cont.CreateUserController)

	// Users Handler
	apiUser.GET("/u", cont.GetUsersController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.GET("/u/:id", cont.GetUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/u/:id", cont.UpdateUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.DELETE("/u/:id", cont.DeleteUserController, middleware.JWT([]byte(conf.JWT_KEY)))
}

func RegisterCourseGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewCourseRepository(db)

	svc := service.NewCourseService(repo)

	cont := CourseController{
		service: svc,
	}

	authCourse := e.Group("/course",
		middleware.Logger(),
		middleware.CORS(),
		m.APIKEYMiddleware,
	)

	authCourse.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	//authcourse handler
	authCourse.POST("/create", cont.CreateCourse)
	authCourse.PUT("/edit/:id", cont.EditCourse)
	authCourse.DELETE("/delete/:id", cont.DeleteCourse)

	courseGroup := e.Group("/course",
		middleware.Logger(),
		middleware.CORS(),
	)
	//course handler
	courseGroup.GET("/all", cont.GetCourses)
	courseGroup.GET("/:id", cont.GetCourse)
}
