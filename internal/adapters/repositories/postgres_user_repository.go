package repositories

import (
	"database/sql"
	"fmt"

	ormentities "github.com/BrunoPolaski/auth-service/internal/adapters/repositories/orm_entities"
	"github.com/BrunoPolaski/auth-service/internal/core/entities"
	"github.com/BrunoPolaski/go-rest-err/rest_err"
)

type userRepository struct {
	database *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) FindUserByEmail(username string) (*entities.User, *rest_err.RestErr) {
	userORM := &ormentities.UserORM{}

	stmt, _ := ur.database.Prepare("SELECT * FROM users WHERE email = $1")
	defer stmt.Close()

	err := stmt.QueryRow(username).Scan(&userORM.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("User not found")
		}
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error while trying to find user: %v", err.Error()))
	}

	user, err := userORM.ToDomain()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return user, nil
}

func (ur *userRepository) SignUp(u *entities.User) *rest_err.RestErr {
	query := `
		INSERT INTO users(email, password, user_type, tenant_id)
		VALUES($1, $2, $3, $4)
	`
	res, err := ur.database.Exec(query,
		u.Email().Value(),
		u.Password().Value(),
		// add here the other properties
	)

	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return rest_err.NewInternalServerError("user was not created")
	}

	return nil
}
