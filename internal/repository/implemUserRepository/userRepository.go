package implemUserRepository

import (
	"eventPlanner/internal/repository"
	"github.com/jackc/pgx/v4"
)

type PostgresUserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) repository.UserRepository {
	return &PostgresUserRepository{conn: conn}
}
