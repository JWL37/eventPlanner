package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log/slog"
)

func ConnectDB(log *slog.Logger, dsn string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Error("Unable to connect to database: %v\n", err)
		panic("Unable to connect to database")
	}
	return conn
}
