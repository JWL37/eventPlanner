package implemContactService

import (
	"eventPlanner/internal/repositories"
	"eventPlanner/internal/services"
)

type contactService struct {
	repo repositories.ContactRepository
}

func NewContactService(repo repositories.ContactRepository) services.ContactService {
	return &contactService{repo: repo}
}
