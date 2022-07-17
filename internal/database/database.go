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
	DB.Create(&model.Role{
		RoleID:      2,
		RoleName:    "User",
		Description: "User",
	})

	DB.Create(&model.User{
		RoleID:    1,
		UserID:    1,
		Email:     "Admin",
		Password:  "Admin",
		ProfileID: 0,
	})

	DB.Create(&model.CourseType{
		CourseTypeID: 1,
		CourseType:   "Lifetime",
	})
	DB.Create(&model.CourseType{
		CourseTypeID: 2,
		CourseType:   "Bootcamp",
	})

	DB.Create(&model.CourseCategory{
		CourseCategoryID: 1,
		Category:         "Website",
	})
	DB.Create(&model.CourseCategory{
		CourseCategoryID: 2,
		Category:         "Mobile",
	})

	DB.Create(&model.CourseCategory{
		CourseCategoryID: 99,
		Category:         "Other",
	})

	DB.AutoMigrate(
		&model.User{},
		&model.UserProfile{},
		&model.Role{},
		&model.Course{},
		&model.Material{},
		&model.CourseCategory{},
		&model.CourseType{},
		&model.CourseDetail{},
	)
	return DB
}
