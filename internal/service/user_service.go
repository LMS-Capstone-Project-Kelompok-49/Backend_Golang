package service

import (
	"errors"
	"net/http"

	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/config"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/domain"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/helper"
	"github.com/LMS-Capstone-Project-Kelompok-49/Backend-Golang/internal/model"
	"gopkg.in/go-playground/validator.v9"
)

type svcUser struct {
	c    config.Config
	repo domain.UserAdapterRepository
}

func (s *svcUser) CreateUserService(user model.User) (error, int) {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err, http.StatusBadRequest
	}
	usr, _ := s.repo.GetOneByEmail(user.Email)
	if usr.Email == user.Email {
		err = errors.New("Email already registered!")
		return err, http.StatusForbidden
	}
	return s.repo.CreateUsers(user), http.StatusOK
	// return nil, http.StatusOK
}

func (s *svcUser) UpdateUserService(id int, user model.User) error {
	return s.repo.UpdateOneByID(id, user)
}

func (s *svcUser) GetAllUsersService() []model.User {
	return s.repo.GetAll()
}

func (s *svcUser) GetUserByID(id int) (model.User, error) {
	return s.repo.GetOneByID(id)
}

func (s *svcUser) LoginUser(email, password string) (string, int) {
	if len(email) == 0 {
		return "Email atau Password tidak boleh kosong", http.StatusBadRequest
	}
	user, _ := s.repo.GetOneByEmail(email)

	if (user.Password != password) && (user.Email != email) {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(int(user.UserID), user.Email, user.RoleID, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func (s *svcUser) DeleteByID(id int) error {
	return s.repo.DeleteByID(id)
}

func NewServiceUser(repo domain.UserAdapterRepository, c config.Config) domain.UserAdapterService {
	return &svcUser{
		repo: repo,
		c:    c,
	}
}
