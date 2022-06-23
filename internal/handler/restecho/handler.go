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
	apiUser := e.Group("/app",
		middleware.Logger(),
		middleware.CORS(),
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

	authCourse := e.Group("/app",
		middleware.Logger(),
		middleware.CORS(),
	)

	authCourse.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	//authcourse handler
	authCourse.POST("/c/create", cont.CreateCourse)
	authCourse.PUT("/c/edit/:id", cont.EditCourse)
	authCourse.DELETE("/c/delete/:id", cont.DeleteCourse)

	courseGroup := e.Group("/app",
		middleware.Logger(),
		middleware.CORS(),
	)
	//course handler
	courseGroup.GET("/c/all", cont.GetCourses)
	courseGroup.GET("/c/:id", cont.GetCourse)
}

func RegisterRoleGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewRoleRepository(db)

	svc := service.NewRoleService(repo)

	cont := RoleController{
		service: svc,
	}

	authRole := e.Group("/app",
		middleware.Logger(),
		middleware.CORS(),
	)

	authRole.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	//authRole handler
	authRole.POST("/r/create", cont.CreateRole)
	authRole.PUT("/r/edit/:id", cont.EditRole)
	authRole.DELETE("/r/delete/:id", cont.DeleteRole)

	roleGroup := e.Group("/app",
		middleware.Logger(),
		middleware.CORS(),
	)
	//role handler
	roleGroup.GET("/r/:id", cont.GetRole)

}

func RegisterCourseCategoryGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewCourseCategoryRepository(db)

	svc := service.NewCourseCategoryService(repo)

	cont := CourseCategoryController{
		service: svc,
	}

	authRole := e.Group("/app",
		middleware.Logger(),
		middleware.CORS(),
	)

	authRole.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	//authCCategory handler
	authRole.POST("/cc/create", cont.CreateCourseCategory)

	roleGroup := e.Group("/app",
		middleware.Logger(),
		middleware.CORS(),
	)
	//ccategory handler
	roleGroup.GET("/cc", cont.GetCoursesCategory)

}

func RegisterTypeCategoryGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewTypeCategoryRepository(db)

	svc := service.NewTypeCategoryService(repo)

	cont := TypeCategoryController{
		service: svc,
	}

	// authRole := e.Group("/app",
	// 	middleware.Logger(),
	// 	middleware.CORS(),
	// )

	// authRole.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	// //authCCategory handler
	// authRole.POST("/cc/create", cont.CreateCourseCategory)

	roleGroup := e.Group("/app",
		middleware.Logger(),
		middleware.CORS(),
	)
	//ccategory handler
	roleGroup.GET("/tc", cont.GetTypeCourse)

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
		m.APIKEYMiddleware,
	)

	authMaterial.Use(middleware.JWT([]byte(conf.JWT_KEY)))

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
