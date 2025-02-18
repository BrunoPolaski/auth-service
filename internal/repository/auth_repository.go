package repository

import "database/sql"

type AuthRepository interface {
	FindUserByEmail(username, password string) error
}

type authRepository struct {
	database *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		database: db,
	}
}

func (ar authRepository) SignIn(username, password string) error {
	return nil
}
