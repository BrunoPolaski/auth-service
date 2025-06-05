package services

import (
	"github.com/BrunoPolaski/auth-service/internal/adapters/repositories/mysql"
	"github.com/BrunoPolaski/auth-service/internal/config/crypto"
)

type AuthService interface {
	SignIn(username, password string) (string, *rest_err.RestErr)
}

type authService struct {
	authRepository mysql.AuthRepository
	crypto         crypto.Crypto
}

func NewAuthService(authRepository mysql.AuthRepository, crypto crypto.Crypto) AuthService {
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
