package routes

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/login-service/internal/config/database"
	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/BrunoPolaski/login-service/internal/controller"
	"github.com/BrunoPolaski/login-service/internal/domain/service"
	"github.com/BrunoPolaski/login-service/internal/repository"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("Failed to initialize routes: %v", r))
		}
	}()

	logger.Info("Initializing routes")
	r := mux.NewRouter()

	var databaseAdapter database.Database = &database.PostgresAdapter{}

	authRepository := repository.NewAuthRepository(databaseAdapter)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signin", authController.SignIn).Methods(http.MethodPost)
	}

	return r
}
