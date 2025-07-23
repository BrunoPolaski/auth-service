package repositories

import (
	"github.com/BrunoPolaski/auth-service/internal/core/entities"
	"github.com/BrunoPolaski/go-rest-err/rest_err"
)

type AuthRepository interface {
	FindUserByEmail(username, password string) (*entities.User, *rest_err.RestErr)
}
