package router

import (
	"auth/internal/service"
	"github.com/labstack/echo/v4"
)

func InitRouter(authService service.AuthService, e *echo.Echo) {
	e.POST("/auth/register", authService.Register)
	e.POST("/auth/login", authService.Login)

}
