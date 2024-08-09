package router

import (
	"eventPlanner/internal/services"
	"github.com/labstack/echo/v4"
)

func InitRouter(userService services.UserService, eventService services.EventService, contactService services.ContactService, e *echo.Echo) {
	e.POST("/auth/register", userService.Register)
	e.POST("/auth/login", userService.Login)

	e.POST("/events/create/:id", eventService.CreateEvent)
	e.GET("/events/:id", eventService.GetAllEvents)

	e.GET("/contacts", contactService.GetAllContacts)
}
