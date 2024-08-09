package implemEventService

import (
	"eventPlanner/internal/repository"
	"eventPlanner/internal/services"
)

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) services.EventService {
	return &eventService{
		repo: repo,
	}
}
