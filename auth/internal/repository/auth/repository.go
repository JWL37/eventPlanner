package auth

import (
	"auth/internal/repository"
	"github.com/jackc/pgx/v4"
)

type PostgresAuthRepository struct {
	conn *pgx.Conn
}

func NewAuthRepository(conn *pgx.Conn) repository.AuthRepository {
	return &PostgresAuthRepository{conn: conn}
}
