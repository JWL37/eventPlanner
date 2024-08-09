package repository

import "eventPlanner/internal/models"

type UserRepository interface {
	CreateUser(user models.User) (int, error)
	FindUser(username string) (*models.User, error)
}
