package ormentities

import (
	"github.com/BrunoPolaski/auth-service/internal/core/entities"
)

type ApiKeyORM struct {
	Id          int64
	Description string
	TenantId    int64
}

func (u *ApiKeyORM) ToDomain() (*entities.ApiKey, error) {
	return entities.NewApiKey(
		u.Id,
		u.Description,
		u.TenantId,
	), nil
}
