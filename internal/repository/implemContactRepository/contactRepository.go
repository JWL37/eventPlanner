package implemContactRepository

import (
	"eventPlanner/internal/repository"
	"github.com/jackc/pgx/v4"
)

type contactRepository struct {
	conn *pgx.Conn
}

func NewContactRepository(conn *pgx.Conn) repository.ContactRepository {
	return &contactRepository{conn: conn}
}
