package app

import (
	"auth/internal/config"
	"auth/internal/database"
	repAuth "auth/internal/repository/auth"
	"auth/internal/router"
	serviceAuth "auth/internal/service/auth"
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

	authRepo := repAuth.NewAuthRepository(conn)
	authService := serviceAuth.NewAuthService(authRepo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.InitRouter(authService, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
