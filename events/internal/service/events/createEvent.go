package events

import (
	"events/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

func (s *eventService) CreateEvent(c echo.Context) error {
	organizerID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	event := new(models.Event)
	if err = c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	event, err = s.repo.CreateEvent(organizerID, *event)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	for _, el := range event.ListMembers {
		memberID, _ := strconv.ParseInt(el, 10, 64)
		if err = s.repo.AddParticipant(memberID, event); err != nil {
			log.Error(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusCreated, event)
}
