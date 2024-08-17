package implemEventService

import (
	"eventPlanner/internal/repositories"
	"eventPlanner/internal/services"
)

type eventService struct {
	repo repositories.EventRepository
}

func NewEventService(repo repositories.EventRepository) services.EventService {
	return &eventService{
		repo: repo,
	}
}
