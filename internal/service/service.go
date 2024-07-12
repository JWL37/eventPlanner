package service

import "github.com/labstack/echo/v4"

type UserService interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type EventService interface {
	CreateEvent(c echo.Context) error
	GetAllEvents(c echo.Context) error
}

type ContactService interface {
	GetAllContacts(c echo.Context) error
	SelectContacts(c echo.Context) error
}
