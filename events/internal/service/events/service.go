package events

import (
	"events/internal/repository"
	"events/internal/service"
)

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) service.EventService {
	return &eventService{
		repo: repo,
	}
}
