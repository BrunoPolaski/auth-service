package repositories

import (
	"github.com/BrunoPolaski/auth-service/internal/infra/thirdparty/database"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
)

type AuthRepository interface {
	FindUserByEmail(username, password string) *rest_err.RestErr
}

type authRepository struct {
	database database.Database
}

func NewAuthRepository(db database.Database) AuthRepository {
	return &authRepository{
		database: db,
	}
}

func (ar authRepository) FindUserByEmail(username, password string) *rest_err.RestErr {
	return nil
}
