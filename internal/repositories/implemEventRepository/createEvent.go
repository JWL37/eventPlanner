package implemEventRepository

import (
	"context"
	"eventPlanner/internal/models"
)

const (
	createEventQuery = `
        INSERT INTO events (organizer_id, name_event, shape, place, begin_time, duration)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
	createParticipationQuery = `
        INSERT INTO participation (user_id, event_id)
        VALUES ($1, $2)`
)

func (r *eventRepository) CreateEvent(organizerID int64, event models.Event) (*models.Event, error) {
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	err = tx.QueryRow(context.Background(), createEventQuery, organizerID, event.NameEvent, event.Shape, event.Place, event.BeginTime, event.Duration).Scan(&event.ID)
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(context.Background(), createParticipationQuery, organizerID, event.ID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) AddParticipant(userID int64, event *models.Event) error {
	_, err := r.conn.Exec(context.Background(), createParticipationQuery, userID, event.ID)
	if err != nil {
		return err
	}
	return nil
}
