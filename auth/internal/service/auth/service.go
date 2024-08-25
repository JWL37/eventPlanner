package auth

import (
	"auth/internal/repository"
	"auth/internal/service"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) service.AuthService {
	return &AuthService{
		repo: repo,
	}
}
