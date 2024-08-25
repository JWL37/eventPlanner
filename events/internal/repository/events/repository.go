package events

import (
	"events/internal/repository"
	"github.com/jackc/pgx/v4"
)

type eventRepository struct {
	conn *pgx.Conn
}

func NewEventRepository(conn *pgx.Conn) repository.EventRepository {
	return &eventRepository{conn: conn}
}
