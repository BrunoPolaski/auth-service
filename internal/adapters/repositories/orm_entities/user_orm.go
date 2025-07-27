package ormentities

import (
	"time"

	"github.com/BrunoPolaski/auth-service/internal/core/entities"
	valueobjects "github.com/BrunoPolaski/auth-service/internal/core/value_objects"
	"github.com/BrunoPolaski/go-rest-err/rest_err"
)

type UserORM struct {
	Id                  int64     `db:"id"`
	Email               string    `db:"email"`
	Password            string    `db:"password"`
	NeedsPasswordChange bool      `db:"needs_password_change"`
	CreatedAt           time.Time `db:"created_at"`
	IsActive            bool      `db:"is_active"`
}

func (u *UserORM) ToDomain() (*entities.User, error) {
	pwd, err := valueobjects.NewPassword(u.Password)
	if err != nil {
		return &entities.User{}, rest_err.NewBadRequestError(err.Error())
	}

	email, err := valueobjects.NewEmail(u.Email)
	if err != nil {
		return &entities.User{}, rest_err.NewBadRequestError(err.Error())
	}

	return entities.NewUserWithFields(
		u.Id,
		email,
		pwd,
		u.CreatedAt,
		u.IsActive,
		u.NeedsPasswordChange,
	), nil
}
