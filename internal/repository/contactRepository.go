package repository

import "eventPlanner/internal/models"

type ContactRepository interface {
	GetAllContacts() ([]models.Contact, error)
}
