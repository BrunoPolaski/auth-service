package routes

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/login-service/internal/domain/service"
	"github.com/BrunoPolaski/login-service/internal/interfaces/http/controllers"
	"github.com/BrunoPolaski/login-service/internal/interfaces/http/middlewares"
	"github.com/BrunoPolaski/login-service/internal/repository"
	"github.com/BrunoPolaski/login-service/internal/thirdparty/crypto"
	"github.com/BrunoPolaski/login-service/internal/thirdparty/database"
	"github.com/BrunoPolaski/login-service/internal/thirdparty/jwt"
	"github.com/BrunoPolaski/login-service/internal/thirdparty/logger"
)

func Init() http.Handler {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("Failed to initialize routes: %v", r))
		}
	}()

	logger.Info("Initializing routes")

	var (
		bcryptAdapter   crypto.Crypto     = crypto.NewBcryptAdapter()
		postgresAdapter database.Database = database.NewPostgresAdapter()
		jwtAdapter      jwt.JWT           = jwt.NewJWTAdapter()

		loggingMiddleware    middlewares.Middleware = middlewares.LoggingMiddleware
		bearerAuthMiddleware middlewares.Middleware = middlewares.BearerMiddleware(jwtAdapter)
		basicAuthMiddleware  middlewares.Middleware = middlewares.BasicMiddleware
	)

	authRepository := repository.NewAuthRepository(postgresAdapter)
	authService := service.NewAuthService(
		authRepository,
		bcryptAdapter,
	)
	authController := controllers.NewAuthController(authService)

	r := http.NewServeMux()
	{
		r.Handle("/auth/signin", middlewares.CreateStack([]middlewares.Middleware{basicAuthMiddleware}, http.HandlerFunc(authController.SignIn)))
		r.Handle("/auth/refresh-token", middlewares.CreateStack([]middlewares.Middleware{bearerAuthMiddleware}, http.HandlerFunc(authController.RefreshToken)))

		r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.Handle("/", http.NotFoundHandler())
	}

	return loggingMiddleware(r)
}
