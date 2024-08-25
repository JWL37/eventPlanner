package repository

import "contacts/internal/models"

type ContactRepository interface {
	GetAllContacts() ([]models.Contact, error)
}
