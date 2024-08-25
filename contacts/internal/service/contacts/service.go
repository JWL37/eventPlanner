package contacts

import (
	"contacts/internal/repository"
	"contacts/internal/service"
)

type contactService struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) service.ContactService {
	return &contactService{repo: repo}
}
