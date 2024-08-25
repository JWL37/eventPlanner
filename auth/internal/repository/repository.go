package repository

import "auth/internal/models"

type AuthRepository interface {
	CreateUser(user models.User) (int, error)
	FindUser(username, password string) (*models.User, error)
}
