package app

import (
	"context"
	"events/internal/config"
	"events/internal/database"
	repEvents "events/internal/repository/events"
	"events/internal/router"
	serviceEvents "events/internal/service/events"
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

	eventRepo := repEvents.NewEventRepository(conn)
	eventService := serviceEvents.NewEventService(eventRepo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.InitRouter(eventService, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
