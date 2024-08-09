package implemContactService

import (
	"eventPlanner/internal/repository"
	"eventPlanner/internal/services"
)

type contactService struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) services.ContactService {
	return &contactService{repo: repo}
}
