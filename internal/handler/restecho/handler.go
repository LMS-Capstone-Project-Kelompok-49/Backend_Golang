package restecho

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/config"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/database"
	assignment "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho/assignment"
	course "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho/course"
	course_cat "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho/course_cat"
	detail "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho/course_detail"
	course_t "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho/course_type"
	material "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho/material"
	profile "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho/profile"
	rating "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho/rating"
	role "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho/role"
	user "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho/user"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/repository"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/service"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepository(db)

	svc := service.NewServiceUser(repo, conf)

	cont := user.EchoController{
		Svc: svc,
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
	prepo := repository.NewProfileRepository(db)
	catRepo := repository.NewCourseCategoryRepository(db)
	eRepo := repository.NewEnrollmentRepository(db)
	rRepo := repository.NewRatingRepository(db)

	svc := service.NewCourseService(repo, prepo)
	catSvc := service.NewCourseCategoryService(catRepo)
	eSvc := service.NewEnrollmentService(eRepo, repo)
	rSvc := service.NewRatingService(rRepo, eRepo)

	cont := course.CourseController{
		Service:    svc,
		CatService: catSvc,
		RService:   rSvc,
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

	//------------------ mentor
	courseMentor := e.Group("/mentor")
	courseMentor.GET("/course/:course_id", cont.GetCourseDash)

	//------------------ user

	eCont := course.UserCourseController{
		Service:    svc,
		EService:   eSvc,
		CatService: catSvc,
		RService:   rSvc,
	}

	enrCourse := e.Group("/enrollment",
		middleware.Logger(),
		middleware.CORS(),
	)

	enrCourse.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	enrCourse.POST("/join/:course_id", eCont.JoinCourse)
	enrCourse.PUT("/update/:course_id", eCont.Update)

	usrDash := e.Group("/user",
		middleware.Logger(),
		middleware.CORS(),
	)
	usrDash.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	usrDash.GET("/course", eCont.GetAll)
	usrDash.GET("/course/:course_id", eCont.GetByID)
	usrDash.GET("/dashboard", eCont.Dashboard)
}

func RegisterRoleGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewRoleRepository(db)

	svc := service.NewRoleService(repo)

	cont := role.RoleController{
		Service: svc,
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

	cont := course_cat.CourseCategoryController{
		Service: svc,
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

	cont := course_t.TypeCategoryController{
		Service: svc,
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
	prepo := repository.NewProfileRepository(db)

	cRepo := repository.NewCourseRepository(db)

	svc := service.NewMaterialService(repo)
	cSvc := service.NewCourseService(cRepo, prepo)

	cont := material.MaterialController{
		Service:       svc,
		CourseService: cSvc,
	}

	authMaterial := e.Group("/material",
		middleware.Logger(),
		middleware.CORS(),
	)

	authMaterial.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	//authMaterial handler
	authMaterial.POST("/create/:courseid", cont.CreateMaterial)
	authMaterial.PUT("/edit/:id", cont.EditMaterial)
	authMaterial.DELETE("/delete/:id", cont.DeleteMaterial)
	//--

	materialGroup := e.Group("/material",
		middleware.Logger(),
		middleware.CORS(),
	)

	//material handler
	materialGroup.GET("/course/:courseid", cont.GetMaterials)
	materialGroup.GET("/:id", cont.GetMaterial)
	//--
}

func RegisterProfileGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewProfileRepository(db)

	svc := service.NewProfileService(repo)

	cont := profile.ProfileController{
		Service: svc,
	}

	authProfile := e.Group("/user/profile",
		middleware.Logger(),
		middleware.CORS(),
	)

	authProfile.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	//authMaterial handler
	authProfile.POST("/create", cont.CreateProfile)
	authProfile.PUT("/edit", cont.EditProfile)
	// authMaterial.DELETE("/delete", cont.)
	authProfile.GET("", cont.GetProfile)
	//--

}

func RegisterCourseDetailAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewCourseDetailRepository(db)
	svc := service.NewCourseDetailService(repo)

	cont := detail.CourseDetailController{
		Service: svc,
	}

	detailGroup := e.Group("/course",
		middleware.Logger(),
		middleware.CORS(),
	)

	detailGroup.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	detailGroup.PUT("/detail/edit/:course_id", cont.EditDetail)
	detailGroup.GET("/detail/:course_id", cont.GetByID)
}

func RegisterRatingAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewRatingRepository(db)
	prepo := repository.NewProfileRepository(db)
	cRepo := repository.NewCourseRepository(db)
	eRepo := repository.NewEnrollmentRepository(db)
	svc := service.NewRatingService(repo, eRepo)
	cSvc := service.NewCourseService(cRepo, prepo)

	cont := rating.RatingController{
		Service:       svc,
		CourseService: cSvc,
	}

	ratingAuth := e.Group("/course/rating",
		middleware.Logger(),
		middleware.CORS(),
	)

	ratingAuth.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	ratingAuth.POST("/create/:course_id", cont.Create)
	ratingAuth.PUT("/edit/:id", cont.Update)
	ratingAuth.DELETE("/delete/:id", cont.Delete)

	ratingGroup := e.Group("/course/rating",
		middleware.Logger(),
		middleware.CORS(),
	)

	ratingGroup.GET("/:course_id", cont.GetByCourse)
	ratingGroup.GET("/user/:id", cont.GetByID)
}

func RegisterAssignmentAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)
	repo := repository.NewAssignmentMentorRepository(db)
	uRepo := repository.NewAssignmentUseerRepository(db)
	cRepo := repository.NewCourseRepository(db)
	pRepo := repository.NewProfileRepository(db)
	svc := service.NewAssignmentMentorService(repo)
	uSvc := service.NewAssignmentUserService(uRepo)
	cSvc := service.NewCourseService(cRepo, pRepo)

	cont := assignment.AssignmentMentorController{
		Service: svc,
	}

	uCont := assignment.AssignmentUserController{
		Service:  uSvc,
		MService: svc,
		CService: cSvc,
	}

	//----mentor

	assignmentMentor := e.Group("/mentor/assignment",
		middleware.Logger(),
		middleware.CORS(),
	)

	assignmentMentor.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	assignmentMentor.POST("/create/:course_id", cont.CreateAssignment)
	assignmentMentor.PUT("/update/:id", cont.EditAssignment)
	assignmentMentor.DELETE("/delete/:id", cont.DeleteAssignment)
	assignmentMentor.GET("/:id", cont.GetAssignment)

	//----user
	assignmentUser := e.Group("/user/assignment",
		middleware.Logger(),
		middleware.CORS(),
	)
	assignmentUser.Use(middleware.JWT([]byte(conf.JWT_KEY)))

	assignmentUser.GET("/:id", uCont.GetAssignmentByID)
	assignmentUser.POST("/create/:am_id", uCont.CreateAssignment)
}
