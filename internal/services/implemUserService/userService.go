package implemUserService

import (
	"eventPlanner/internal/repository"
	"eventPlanner/internal/services"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) services.UserService {
	return &userService{
		repo: repo,
	}
}
