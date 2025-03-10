package routes

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/login-service/internal/config/crypto"
	"github.com/BrunoPolaski/login-service/internal/config/database"
	"github.com/BrunoPolaski/login-service/internal/config/jwt"
	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/BrunoPolaski/login-service/internal/controller"
	"github.com/BrunoPolaski/login-service/internal/controller/middleware"
	"github.com/BrunoPolaski/login-service/internal/domain/service"
	"github.com/BrunoPolaski/login-service/internal/repository"
)

func Init() http.Handler {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("Failed to initialize routes: %v", r))
		}
	}()

	logger.Info("Initializing routes")

	var (
		cryptoAdapter   crypto.Crypto     = crypto.NewBcryptAdapter()
		databaseAdapter database.Database = database.NewPostgresAdapter()
		jwtAdapter      jwt.JWT           = jwt.NewJWTAdapter()

		loggingMiddleware middleware.Middleware = middleware.LoggingMiddleware
		bearerMiddleware  middleware.Middleware = middleware.BearerMiddleware(jwtAdapter)
	)

	authRepository := repository.NewAuthRepository(databaseAdapter)
	authService := service.NewAuthService(
		authRepository,
		cryptoAdapter,
	)
	authController := controller.NewAuthController(authService)

	r := http.NewServeMux()
	{
		r.HandleFunc("/auth/signin", authController.SignIn)

		r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.Handle("/", http.NotFoundHandler())
	}

	return loggingMiddleware(r)
}
