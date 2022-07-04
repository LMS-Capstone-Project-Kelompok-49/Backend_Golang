package main

import (
	"net/http"
	"os"

	conf "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/config"
	rest "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho"
	"github.com/labstack/echo/v4"

	_ "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/docs" // docs is generated by Swag CLI, you have to import it.
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description Sample API
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	config := conf.InitConfiguration()
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Welcome! You can check documentation at http://3.95.181.246/swagger/index.html !",
		})
	})

	os.Mkdir("temp", 0755)

	rest.RegisterUserGroupAPI(e, config)
	rest.RegisterCourseGroupAPI(e, config)
	rest.RegisterRoleGroupAPI(e, config)
	rest.RegisterMaterialGroupAPI(e, config)
	rest.RegisterCourseCategoryGroupAPI(e, config)
	rest.RegisterTypeCategoryGroupAPI(e, config)

	e.Logger.Fatal(e.Start(":8000"))
}
