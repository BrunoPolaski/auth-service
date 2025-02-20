package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/login-service/internal/repository"
)

type AuthService interface {
	SignIn(username, password string) (string, *rest_err.RestErr)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (as *authService) SignIn(username, password string) (string, *rest_err.RestErr) {
	return "", rest_err.NewNotFoundError("not implemented")
}
