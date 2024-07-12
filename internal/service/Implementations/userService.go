package Implementations

import (
	"eventPlanner/internal/models"
	"eventPlanner/internal/repository"
	"eventPlanner/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) service.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Register(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	id, err := s.repo.CreateUser(*u)
	if err != nil {
		return c.JSON(http.StatusConflict, err.Error())
	}
	return c.JSON(http.StatusCreated, id)
}

func (s *userService) Login(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	user, err := s.repo.FindUser(u.Name)
	if err != nil || user.Password != u.Password {
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}

	return c.JSON(http.StatusOK, user.ID)
}
