package repositories

import (
	"github.com/BrunoPolaski/auth-service/internal/core/entities"
	"github.com/BrunoPolaski/go-rest-err/rest_err"
)

type UserRepository interface {
	FindUserByEmail(username string) (*entities.User, *rest_err.RestErr)
}

type TokenRepository interface {
	FindById(token string) *entities.User
}
