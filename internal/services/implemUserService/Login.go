package implemUserService

import (
	"eventPlanner/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *userService) Login(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	user, err := s.repo.FindUser(u.Name, u.Password)
	if err != nil || user.Password != u.Password {
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}

	return c.JSON(http.StatusOK, user.ID)
}
