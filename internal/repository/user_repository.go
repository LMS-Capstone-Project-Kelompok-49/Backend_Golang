package repository

import (
	"fmt"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gorm.io/gorm"
)

type repositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *repositoryMysqlLayer) CreateUsers(user model.User) error {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAll() []model.User {
	users := []model.User{}
	r.DB.Preload("Profile").Find(&users)

	return users
}

func (r *repositoryMysqlLayer) GetOneByID(id int) (user model.User, err error) {
	res := r.DB.Where("user_id = ?", id).Preload("Profile").Preload("Enrollment").Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repositoryMysqlLayer) GetOneByEmail(email string) (user model.User, err error) {
	res := r.DB.Where("email = ?", email).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateOneByID(id int, user model.User) error {
	res := r.DB.Where("user_id = ?", id).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteByID(id int) error {
	res := r.DB.Delete(&model.User{
		UserID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlRepository(db *gorm.DB) domain.UserAdapterRepository {
	return &repositoryMysqlLayer{
		DB: db,
	}
}
