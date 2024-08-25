package contacts

import (
	"contacts/internal/models"
	"context"
)

const (
	getAllContactsQuery = `SELECT id, username, email FROM users`
)

func (r *contactRepository) GetAllContacts() ([]models.Contact, error) {
	rows, err := r.conn.Query(context.Background(), getAllContactsQuery)
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
