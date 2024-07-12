package Implementations

import (
	"context"
	"eventPlanner/internal/models"
	"eventPlanner/internal/repository"
	"github.com/jackc/pgx/v4"
)

type eventRepository struct {
	conn *pgx.Conn
}

func NewEventRepository(conn *pgx.Conn) repository.EventRepository {
	return &eventRepository{conn: conn}
}

func (r *eventRepository) CreateEvent(userID int64, event models.Event) error {
	query := `
        INSERT INTO events (user_id, name_event, shape, place, begin_time, duration)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
	err := r.conn.QueryRow(context.Background(), query, userID, event.NameEvent, event.Shape, event.Place, event.BeginTime, event.Duration).Scan(&event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *eventRepository) GetAllEvents(userID int64) ([]models.Event, error) {
	query := `
        SELECT id, name_event, shape, place, begin_time, duration
        FROM events
        WHERE user_id = $1
    `
	rows, err := r.conn.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.NameEvent, &event.Shape, &event.Place, &event.BeginTime, &event.Duration)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
