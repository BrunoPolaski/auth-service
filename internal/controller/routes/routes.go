package routes

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/login-service/internal/config/crypto"
	"github.com/BrunoPolaski/login-service/internal/config/database"
	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/BrunoPolaski/login-service/internal/controller"
	"github.com/BrunoPolaski/login-service/internal/domain/service"
	"github.com/BrunoPolaski/login-service/internal/repository"
)

func Init() *http.ServeMux {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("Failed to initialize routes: %v", r))
		}
	}()

	logger.Info("Initializing routes")
	r := http.NewServeMux()

	var databaseAdapter database.Database = database.NewPostgresAdapter()
	var cryptoAdapter crypto.Crypto = crypto.NewBcryptAdapter()

	authRepository := repository.NewAuthRepository(databaseAdapter)
	authService := service.NewAuthService(
		authRepository,
		cryptoAdapter,
	)
	authController := controller.NewAuthController(authService)

	r.HandleFunc("POST /auth", authController.SignIn)

	return r
}
