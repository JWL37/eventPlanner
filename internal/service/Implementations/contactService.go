package Implementations

import (
	"eventPlanner/internal/repository"
	"eventPlanner/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type contactService struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) service.ContactService {
	return &contactService{repo: repo}
}

func (s *contactService) GetAllContacts(c echo.Context) error {
	contacts, err := s.repo.GetAllContacts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"contacts": contacts})
}

func (s *contactService) SelectContacts(c echo.Context) error {
	var request struct {
		ContactIDs []string `json:"contact_ids"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest,  "Invalid request")
	}

	contacts, err := s.repo.SelectContacts(request.ContactIDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,  err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success":           true,
		"selected_contacts": contacts,
	})
}
