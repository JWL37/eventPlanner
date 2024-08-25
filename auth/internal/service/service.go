package service

import "github.com/labstack/echo/v4"

type AuthService interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}
