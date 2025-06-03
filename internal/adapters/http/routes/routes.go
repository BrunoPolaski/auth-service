package routes

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/auth-service/internal/adapters/http/controllers"
	"github.com/BrunoPolaski/auth-service/internal/adapters/repositories/mysql"
	"github.com/BrunoPolaski/auth-service/internal/adapters/services"
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

	authRepository := mysql.NewAuthRepository(conn)
	authService := services.NewAuthService(
		authRepository,
		cryptoAdapter,
	)
	authController := controllers.NewAuthController(authService)

	r.HandleFunc("POST /auth", authController.SignIn)

	return r
}
