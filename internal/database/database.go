package database

import (
	"fmt"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/config"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Config) *gorm.DB {

	conectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(conectionString))
	if err != nil {
		fmt.Println("error open conection : ", err)
	}

	DB.Create(&model.Role{
		RoleID:      1,
		RoleName:    "Admin",
		Description: "Admin",
	})

	var modela = []model.CourseType{{CourseTypeID: 1, CourseType: "Lifetime"}, {CourseTypeID: 2, CourseType: "Bootcamp"}}
	DB.Create(modela)

	DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Course{},
		&model.CourseCategory{},
		&model.CourseType{},
	)
	return DB
}
