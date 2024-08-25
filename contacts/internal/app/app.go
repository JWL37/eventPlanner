package app

import (
	"contacts/internal/config"
	"contacts/internal/database"
	repContacts "contacts/internal/repository/contacts"
	"contacts/internal/router"
	serviceContacts "contacts/internal/service/contacts"
	"context"
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

	contactRepo := repContacts.NewContactRepository(conn)
	contactService := serviceContacts.NewContactService(contactRepo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.InitRouter(contactService, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
