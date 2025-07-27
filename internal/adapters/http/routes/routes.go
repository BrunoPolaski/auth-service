package routes

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/auth-service/internal/adapters/http/controllers"
	"github.com/BrunoPolaski/auth-service/internal/adapters/http/middlewares"
	"github.com/BrunoPolaski/auth-service/internal/adapters/repositories"
	"github.com/BrunoPolaski/auth-service/internal/adapters/services"
	"github.com/BrunoPolaski/auth-service/internal/infra/database"
	"github.com/BrunoPolaski/auth-service/internal/infra/logger"
)

func Init() http.Handler {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("Failed to initialize routes: %v", r))
		}
	}()

	r := http.NewServeMux()

	var databaseAdapter database.Database = database.NewPostgresAdapter()

	conn, err := databaseAdapter.Connect()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to database: %v", err))
		return nil
	}

	userRepository := repositories.NewUserRepository(conn)
	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)

	r.Handle("POST /signin", HandlerChain(
		authController.SignIn,
		middlewares.LoggingMiddleware,
	))

	r.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return middlewares.LoggingMiddleware(r)
}
