package routes

import (
	"fmt"
	"net/http"

	controller "github.com/BrunoPolaski/auth-service/internal/adapters/http/controllers"
	repository "github.com/BrunoPolaski/auth-service/internal/adapters/mysql"
	service "github.com/BrunoPolaski/auth-service/internal/adapters/services"
	"github.com/BrunoPolaski/auth-service/internal/config/crypto"
	"github.com/BrunoPolaski/auth-service/internal/config/database"
	"github.com/BrunoPolaski/auth-service/internal/config/logger"
)

func Init() *http.ServeMux {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("Failed to initialize routes: %v", r))
		}
	}()

	r := http.NewServeMux()

	var databaseAdapter database.Database = database.NewPostgresAdapter()
	var cryptoAdapter crypto.Crypto = crypto.NewBcryptAdapter()

	conn, err := databaseAdapter.Connect()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to database: %v", err))
		return nil
	}

	authRepository := repository.NewAuthRepository(conn)
	authService := service.NewAuthService(
		authRepository,
		cryptoAdapter,
	)
	authController := controller.NewAuthController(authService)

	r.HandleFunc("POST /auth", authController.SignIn)

	return r
}
