package repositories

import (
	"github.com/BrunoPolaski/auth-service/internal/core/entities"
	"github.com/BrunoPolaski/go-rest-err/rest_err"
)

type UserRepository interface {
	GetByEmail(username string) (*entities.User, *rest_err.RestErr)
}

type TokenRepository interface {
	GetById(token string) *entities.User
}

type ApiKeyRepository interface {
	GetById(uuid string) (*entities.ApiKey, *rest_err.RestErr)
}
