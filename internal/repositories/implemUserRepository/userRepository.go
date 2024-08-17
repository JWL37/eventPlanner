package implemUserRepository

import (
	"eventPlanner/internal/repositories"
	"github.com/jackc/pgx/v4"
)

type PostgresUserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) repositories.UserRepository {
	return &PostgresUserRepository{conn: conn}
}
