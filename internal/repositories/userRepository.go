package repositories

import "eventPlanner/internal/models"

type UserRepository interface {
	CreateUser(user models.User) (int, error)
	FindUser(username, password string) (*models.User, error)
}
