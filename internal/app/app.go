package app

import (
	"context"
	"eventPlanner/internal/config"
	"eventPlanner/internal/database"
	"eventPlanner/internal/repository/implemContactRepository"
	"eventPlanner/internal/repository/implemEventRepository"
	"eventPlanner/internal/repository/implemUserRepository"
	"eventPlanner/internal/router"
	"eventPlanner/internal/services/implemContactService"
	"eventPlanner/internal/services/implemEventService"
	"eventPlanner/internal/services/implemUserService"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func RunApp(cfg config.Config) {

	dsn := cfg.DataBasePath
	conn, err := database.ConnectDB(dsn)
	if err != nil {
		logrus.Fatalf("Unable to connect to database: %v\n", err)
	}
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
