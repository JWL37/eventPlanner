package Implementations

import (
	"context"
	"eventPlanner/internal/models"
	"eventPlanner/internal/repository"
	"github.com/jackc/pgx/v4"
)

type contactRepository struct {
	conn *pgx.Conn
}

func NewContactRepository(conn *pgx.Conn) repository.ContactRepository {
	return &contactRepository{conn: conn}
}

func (r *contactRepository) GetAllContacts() ([]models.Contact, error) {
	rows, err := r.conn.Query(context.Background(), "SELECT id, username, email FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Email)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (r *contactRepository) SelectContacts(contactIDs []string) ([]models.Contact, error) {
	rows, err := r.conn.Query(context.Background(), "SELECT id, username, email FROM contacts WHERE id = ANY($1::int[])", contactIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Email)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}
