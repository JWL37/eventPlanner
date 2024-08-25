package events

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (s *eventService) GetAllEvents(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	events, err := s.repo.GetAllEvents(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, events)
}
