package services

import "github.com/labstack/echo/v4"

type ContactService interface {
	GetAllContacts(c echo.Context) error
}
