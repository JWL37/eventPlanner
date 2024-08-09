package implemEventRepository

import (
	"context"
	"eventPlanner/internal/models"
)

const (
	createEventQuery = `
        INSERT INTO events (user_id, name_event, shape, place, begin_time, duration)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
)

func (r *eventRepository) CreateEvent(userID int64, event models.Event) error {
	err := r.conn.QueryRow(context.Background(), createEventQuery, userID, event.NameEvent, event.Shape, event.Place, event.BeginTime, event.Duration).Scan(&event.ID)
	if err != nil {
		return err
	}
	return nil
}
