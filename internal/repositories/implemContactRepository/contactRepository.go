package implemContactRepository

import (
	"eventPlanner/internal/repositories"
	"github.com/jackc/pgx/v4"
)

type contactRepository struct {
	conn *pgx.Conn
}

func NewContactRepository(conn *pgx.Conn) repositories.ContactRepository {
	return &contactRepository{conn: conn}
}
