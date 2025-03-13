package routes

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/auth-service/internal/application/services"
	"github.com/BrunoPolaski/auth-service/internal/config/logger"
	"github.com/BrunoPolaski/auth-service/internal/infra/repositories"
	"github.com/BrunoPolaski/auth-service/internal/infra/thirdparty/crypto"
	"github.com/BrunoPolaski/auth-service/internal/infra/thirdparty/database"
	"github.com/BrunoPolaski/auth-service/internal/infra/thirdparty/jwt"
	"github.com/BrunoPolaski/auth-service/internal/interfaces/http/controllers"
	"github.com/BrunoPolaski/auth-service/internal/interfaces/http/middlewares"
)

func Init() http.Handler {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("Failed to initialize routes: %v", r))
		}
	}()

	logger.Info("Initializing routes")

	bcryptAdapter := crypto.NewBcryptAdapter()
	postgresAdapter := database.NewPostgresAdapter()
	jwtAdapter := jwt.NewJWTAdapter()

	loggingMiddleware := middlewares.LoggingMiddleware
	bearerAuthMiddleware := middlewares.BearerMiddleware(jwtAdapter)
	basicAuthMiddleware := middlewares.BasicMiddleware

	authController := controllers.NewAuthController(
		services.NewAuthService(
			repositories.NewAuthRepository(postgresAdapter),
			bcryptAdapter,
		),
	)

	r := http.NewServeMux()
	{
		r.Handle("/auth/signin", basicAuthMiddleware(http.HandlerFunc(authController.SignIn)))
		r.Handle("/auth/refresh-token", bearerAuthMiddleware(http.HandlerFunc(authController.RefreshToken)))

		r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.Handle("/", http.NotFoundHandler())
	}

	return loggingMiddleware(r)
}
