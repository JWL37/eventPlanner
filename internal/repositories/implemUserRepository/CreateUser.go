package implemUserRepository

import (
	"context"
	"eventPlanner/internal/models"
	"fmt"
)

const (
	createUseQuery = `INSERT INTO users (username, email, password, status) VALUES ($1, $2, $3, $4) RETURNING id`
)

func (r *PostgresUserRepository) CreateUser(user models.User) (int, error) {
	_, err := r.FindUser(user.Name, user.Password)
	if err.Error() != "user not found" {
		return 0, fmt.Errorf("error creating user: %w", err)
	}
	var userID int
	err = r.conn.QueryRow(context.Background(), createUseQuery, user.Name, user.Email, user.Password, user.Status).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}
	return userID, nil
}
