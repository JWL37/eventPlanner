package repository

import "eventPlanner/internal/models"

type UserRepository interface {
	CreateUser(user models.User) (int, error)
	FindUser(username string) (*models.User, error)
}

type EventRepository interface {
	CreateEvent(userID int64, event models.Event) error
	GetAllEvents(userID int64) ([]models.Event, error)
}

type ContactRepository interface {
	GetAllContacts() ([]models.Contact, error)
	SelectContacts(contactIDs []string) ([]models.Contact, error)
}
