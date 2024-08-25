package auth

import (
	"auth/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *AuthService) Register(c echo.Context) error {
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
