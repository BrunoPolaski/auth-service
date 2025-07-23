package repositories

import (
	"database/sql"
	"fmt"

	"github.com/BrunoPolaski/auth-service/internal/core/entities"
	"github.com/BrunoPolaski/go-rest-err/rest_err"
)

type authRepository struct {
	database *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		database: db,
	}
}

func (ar *authRepository) FindUserByEmail(username, password string) (*entities.User, *rest_err.RestErr) {
	user := &entities.User{}

	stmt, _ := ar.database.Prepare("SELECT * FROM users WHERE email = $1")
	defer stmt.Close()

	err := ar.database.QueryRow(username).Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("User not found")
		}
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error while trying to find user: %v", err.Error()))
	}

	return user, nil
}
