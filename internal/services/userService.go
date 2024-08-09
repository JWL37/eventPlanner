package services

import "github.com/labstack/echo/v4"

type UserService interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}
