package implemUserRepository

import (
	"context"
	"errors"
	"eventPlanner/internal/models"
	"fmt"
	"github.com/jackc/pgx/v4"
)

const (
	//createUseQuery = `INSERT INTO users (username, email, password, status) VALUES ($1, $2, $3, $4) RETURNING id`
	findUserQuery = `SELECT id, username, email, password, status FROM users WHERE username=$1`
)

func (r *PostgresUserRepository) FindUser(username string) (*models.User, error) {
	var user models.User
	err := r.conn.QueryRow(context.Background(), findUserQuery, username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error finding user: %w", err)
	}
	return &user, nil
}
