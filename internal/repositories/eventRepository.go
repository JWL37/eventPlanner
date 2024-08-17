package repositories

import "eventPlanner/internal/models"

type EventRepository interface {
	// CreateEvent создаёт мероприятие для его организатора
	CreateEvent(userID int64, event models.Event) error
	GetAllEvents(userID int64) ([]models.Event, error)
	// CreateEventMembers создаёт одно и тоже мероприятие для каждого участника
	//CreateEventMembers(userID int64, implemEventRepository models.Event) error
}
