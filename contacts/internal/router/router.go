package router

import (
	"contacts/internal/service"
	"github.com/labstack/echo/v4"
)

func InitRouter(contactService service.ContactService, e *echo.Echo) {

	e.GET("/contacts", contactService.GetAllContacts)
}
