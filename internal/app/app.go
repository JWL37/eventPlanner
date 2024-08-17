package app

import (
	"context"
	"eventPlanner/internal/config"
	"eventPlanner/internal/database"
	"eventPlanner/internal/repositories/implemContactRepository"
	"eventPlanner/internal/repositories/implemEventRepository"
	"eventPlanner/internal/repositories/implemUserRepository"
	"eventPlanner/internal/router"
	"eventPlanner/internal/services/implemContactService"
	"eventPlanner/internal/services/implemEventService"
	"eventPlanner/internal/services/implemUserService"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
)

func RunApp(log *slog.Logger, cfg config.Config) {
	log.Info("starting app")
	dsn := cfg.DataBasePath
	conn := database.ConnectDB(log, dsn)
	defer conn.Close(context.Background())

	userRepo := implemUserRepository.NewUserRepository(conn)
	userService := implemUserService.NewUserService(userRepo)

	eventRepo := implemEventRepository.NewEventRepository(conn)
	eventService := implemEventService.NewEventService(eventRepo)

	contactRepo := implemContactRepository.NewContactRepository(conn)
	contactService := implemContactService.NewContactService(contactRepo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.InitRouter(userService, eventService, contactService, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
