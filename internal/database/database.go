package database

import (
    "context"
    "github.com/jackc/pgx/v4"
    "github.com/sirupsen/logrus"
)

func ConnectDB(dsn string) (*pgx.Conn, error) {
    conn, err := pgx.Connect(context.Background(), dsn)
    if err != nil {
        logrus.Fatalf("Unable to connect to database: %v\n", err)
        return nil, err
    }
    return conn, nil
}