package app

import (
	"context"
	"eventPlanner/internal/config"
	"eventPlanner/internal/database"
	"eventPlanner/internal/repository/Implementations"
	"eventPlanner/internal/router"
	Implementations2 "eventPlanner/internal/service/Implementations"
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

	userRepo := Implementations.NewUserRepository(conn)
	eventRepo := Implementations.NewEventRepository(conn)
	contactRepo := Implementations.NewContactRepository(conn)

	userService := Implementations2.NewUserService(userRepo)
	eventService := Implementations2.NewEventService(eventRepo)
	contactService := Implementations2.NewContactService(contactRepo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.InitRouter(userService, eventService, contactService, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
