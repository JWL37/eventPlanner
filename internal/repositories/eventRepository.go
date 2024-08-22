package repositories

import "eventPlanner/internal/models"

type EventRepository interface {
	// CreateEvent создаёт мероприятие для всех участников мероприятия(организатор в том числе)
	CreateEvent(userID int64, event models.Event) (*models.Event, error)
	GetAllEvents(userID int64) ([]models.Event, error)
	AddParticipant(userID int64, event *models.Event) error
}
