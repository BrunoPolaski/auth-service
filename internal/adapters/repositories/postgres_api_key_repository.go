package repositories

import (
	"database/sql"
	"fmt"

	ormentities "github.com/BrunoPolaski/auth-service/internal/adapters/repositories/orm_entities"
	"github.com/BrunoPolaski/auth-service/internal/core/entities"
	"github.com/BrunoPolaski/go-rest-err/rest_err"
)

type apiKeyRepository struct {
	database *sql.DB
}

func NewApiKeyRepository(db *sql.DB) ApiKeyRepository {
	return &apiKeyRepository{
		database: db,
	}
}

func (ur *apiKeyRepository) GetById(uuid string) (*entities.ApiKey, *rest_err.RestErr) {
	apiKeyORM := &ormentities.ApiKeyORM{}

	stmt, _ := ur.database.Prepare("SELECT * FROM apiKeys WHERE uuid = $1")
	defer stmt.Close()

	err := stmt.QueryRow(uuid).Scan(&apiKeyORM.Id, &apiKeyORM.Description, &apiKeyORM.TenantId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("ApiKey not found")
		}
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error while trying to find apiKey: %v", err.Error()))
	}

	apiKey, err := apiKeyORM.ToDomain()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return apiKey, nil
}
