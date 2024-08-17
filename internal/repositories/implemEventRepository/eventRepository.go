package implemEventRepository

import (
	"eventPlanner/internal/repositories"
	"github.com/jackc/pgx/v4"
)

type eventRepository struct {
	conn *pgx.Conn
}

func NewEventRepository(conn *pgx.Conn) repositories.EventRepository {
	return &eventRepository{conn: conn}
}
