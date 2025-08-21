package ormentities

import (
	"github.com/BrunoPolaski/auth-service/internal/core/entities"
)

type TokenORM struct {
	Id     int64
	UserId int64
}

func (u *TokenORM) ToDomain() (*entities.Token, error) {
	return entities.NewToken(
		u.Id,
		u.UserId,
	), nil
}
