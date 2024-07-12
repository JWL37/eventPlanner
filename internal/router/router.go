package router

import (
	"eventPlanner/internal/service"
	"github.com/labstack/echo/v4"
)

func InitRouter(userService service.UserService, eventService service.EventService, contactService service.ContactService, e *echo.Echo) {
	e.POST("/auth/register", userService.Register)
	e.POST("/auth/login", userService.Login)

	e.POST("/events/create/:id", eventService.CreateEvent)
	e.GET("/events/:id", eventService.GetAllEvents)

	e.GET("/contacts", contactService.GetAllContacts)
	e.POST("/contacts/select", contactService.SelectContacts)
}
