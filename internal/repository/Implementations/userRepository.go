package Implementations

import (
	"context"
	"eventPlanner/internal/models"
	"eventPlanner/internal/repository"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type PostgresUserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) repository.UserRepository {
	return &PostgresUserRepository{conn: conn}
}

func (r *PostgresUserRepository) CreateUser(user models.User) (int, error) {
	var userID int
	query := `INSERT INTO users (username, email, password, status) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.conn.QueryRow(context.Background(), query, user.Name, user.Email, user.Password, user.Status).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}
	return userID, nil
}

func (r *PostgresUserRepository) FindUser(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password, status FROM users WHERE username=$1`
	err := r.conn.QueryRow(context.Background(), query, username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error finding user: %w", err)
	}
	return &user, nil
}
