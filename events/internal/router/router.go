package router

import (
	"events/internal/service"
	"github.com/labstack/echo/v4"
)

func InitRouter(eventService service.EventService, e *echo.Echo) {

	e.POST("/events/create/:id", eventService.CreateEvent)
	e.GET("/events/:id", eventService.GetAllEvents)

}
