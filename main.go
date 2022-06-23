package main

import (
	conf "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/config"
	rest "github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/handler/restecho"
	"github.com/labstack/echo/v4"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()

	rest.RegisterUserGroupAPI(e, config)
	rest.RegisterCourseGroupAPI(e, config)
	rest.RegisterRoleGroupAPI(e, config)
	rest.RegisterMaterialGroupAPI(e, config)

	e.Logger.Fatal(e.Start(":8080"))
}
