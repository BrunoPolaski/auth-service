package factory

import (
	"database/sql"

	"github.com/BrunoPolaski/login-service/internal/controller"
	"github.com/BrunoPolaski/login-service/internal/domain/service"
	"github.com/BrunoPolaski/login-service/internal/repository"
)

type ControllerFactory interface {
	GetAuthController() controller.AuthController
}

type controllerFactory struct {
	database *sql.DB
}

func NewControllerFactory(db *sql.DB) *controllerFactory {
	return &controllerFactory{
		database: db,
	}
}

func (cf *controllerFactory) GetAuthController() controller.AuthController {
	return controller.NewAuthController(
		service.NewAuthService(
			repository.NewAuthRepository(
				cf.database,
			),
		),
	)
}
