package contacts

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *contactService) GetAllContacts(c echo.Context) error {
	contacts, err := s.repo.GetAllContacts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, contacts)
}
