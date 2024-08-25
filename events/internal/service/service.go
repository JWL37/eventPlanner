package service

import "github.com/labstack/echo/v4"

type EventService interface {
	CreateEvent(c echo.Context) error
	GetAllEvents(c echo.Context) error
}
