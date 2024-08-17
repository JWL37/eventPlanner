package implemEventService

import (
	"eventPlanner/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

func (s *eventService) CreateEvent(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	event := new(models.Event)
	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	if err := s.repo.CreateEvent(userID, *event); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	for _, el := range event.ListMembers {
		memberID, _ := strconv.ParseInt(el, 10, 64)
		if err := s.repo.CreateEvent(memberID, *event); err != nil {
			log.Error(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusCreated, event)
}
