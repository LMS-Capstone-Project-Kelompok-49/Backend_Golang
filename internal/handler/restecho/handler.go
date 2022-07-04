package restecho

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/config"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/database"
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
	apiUser := e.Group("/api",
		middleware.Logger(),
		middleware.CORS(),
	)
	// Auth Handler
	auth.POST("/login", cont.LoginUserController)
	auth.POST("/register", cont.CreateUserController)

	// Users Handler
	apiUser.GET("/user/all", cont.GetUsersController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.GET("/user/:user_id", cont.GetUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/user/:user_id", cont.UpdateUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.DELETE("/user/:user_id", cont.DeleteUserController, middleware.JWT([]byte(conf.JWT_KEY)))
}

func RegisterCourseGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewCourseRepository(db)

	svc := service.NewCourseService(repo)

	cont := CourseController{
		service: svc,
	}

	authCourse := e.Group("/api",
		middleware.Logger(),
		middleware.CORS(),
	)

	authCourse.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	//authcourse handler
	authCourse.POST("/course/create", cont.CreateCourse)
	authCourse.PUT("/course/edit/:course_id", cont.EditCourse)
	authCourse.DELETE("/course/delete/:course_id", cont.DeleteCourse)

	courseGroup := e.Group("/api",
		middleware.Logger(),
		middleware.CORS(),
	)
	//course handler
	courseGroup.GET("/course/all", cont.GetCourses)
	courseGroup.GET("/course/:course_id", cont.GetCourse)
}

func RegisterRoleGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewRoleRepository(db)

	svc := service.NewRoleService(repo)

	cont := RoleController{
		service: svc,
	}

	authRole := e.Group("/api",
		middleware.Logger(),
		middleware.CORS(),
	)

	authRole.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	//authRole handler
	// authRole.POST("/role/create", cont.CreateRole)
	// authRole.PUT("/role/edit/:id", cont.EditRole)
	// authRole.DELETE("/r/delete/:id", cont.DeleteRole)

	roleGroup := e.Group("/api",
		middleware.Logger(),
		middleware.CORS(),
	)
	//role handler
	roleGroup.GET("/role/:id", cont.GetRole)

}

func RegisterCourseCategoryGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewCourseCategoryRepository(db)

	svc := service.NewCourseCategoryService(repo)

	cont := CourseCategoryController{
		service: svc,
	}

	authRole := e.Group("/api",
		middleware.Logger(),
		middleware.CORS(),
	)

	authRole.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	//authCCategory handler
	authRole.POST("/course_category/create", cont.CreateCourseCategory)

	roleGroup := e.Group("/api",
		middleware.Logger(),
		middleware.CORS(),
	)
	//ccategory handler
	roleGroup.GET("/course_category", cont.GetCoursesCategory)
	roleGroup.GET("/course_category/:course_category_id", cont.GetCourseCategory)

}

func RegisterTypeCategoryGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewTypeCategoryRepository(db)

	svc := service.NewTypeCategoryService(repo)

	cont := TypeCategoryController{
		service: svc,
	}

	roleGroup := e.Group("/api",
		middleware.Logger(),
		middleware.CORS(),
	)
	//ccategory handler
	roleGroup.GET("/type_course", cont.GetTypeCourse)
	roleGroup.GET("/type_course/:type_course_id", cont.GetOneTypeCourse)

}

func RegisterMaterialGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewMaterialRepository(db)

	svc := service.NewMaterialService(repo)

	cont := MaterialController{
		service: svc,
	}

	authMaterial := e.Group("/material",
		middleware.Logger(),
		middleware.CORS(),
	)

	// authMaterial.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	//authMaterial handler
	authMaterial.POST("/create/:courseid", cont.CreateMaterial)
	authMaterial.PUT("/edit/:id", cont.EditMaterial)
	authMaterial.DELETE("/delete/:id", cont.DeleteMaterial)
	//--

	materialGroup := e.Group("/material")

	//material handler
	materialGroup.GET("/course/:courseid", cont.GetMaterials)
	materialGroup.GET("/:id", cont.GetMaterial)
	//--
}
