package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/login-service/internal/config/crypto"
	"github.com/BrunoPolaski/login-service/internal/repository"
)

type AuthService interface {
	SignIn(username, password string) (string, *rest_err.RestErr)
}

type authService struct {
	authRepository repository.AuthRepository
	crypto         crypto.Crypto
}

func NewAuthService(authRepository repository.AuthRepository, crypto crypto.Crypto) AuthService {
	return &authService{
		authRepository: authRepository,
		crypto:         crypto,
	}
}

func (as *authService) SignIn(username, password string) (string, *rest_err.RestErr) {
	hashedPassword, err := as.crypto.EncryptPassword(password)
	if err != nil {
		return "", err
	}

	err = as.crypto.ComparePasswords(hashedPassword, password)
	if err != nil {
		return "", err
	}

	return "token", nil
}
