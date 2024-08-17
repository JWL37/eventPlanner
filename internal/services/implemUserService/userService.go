package implemUserService

import (
	"eventPlanner/internal/repositories"
	"eventPlanner/internal/services"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) services.UserService {
	return &userService{
		repo: repo,
	}
}
