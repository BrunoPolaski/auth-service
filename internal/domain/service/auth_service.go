package service

import "github.com/BrunoPolaski/login-service/internal/repository"

type AuthService interface {
	SignIn(username, password string) (string, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (as *authService) SignIn(username, password string) (string, error)
