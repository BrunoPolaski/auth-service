package routes

import (
	"net/http"

	"github.com/BrunoPolaski/login-service/internal/config/database"
	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/BrunoPolaski/login-service/internal/controller"
	"github.com/BrunoPolaski/login-service/internal/domain/service"
	"github.com/BrunoPolaski/login-service/internal/repository"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	logger.Info("Initializing routes")
	r := mux.NewRouter()

	databaseAdapter := database.PostgresAdapter{}
	conn, err := databaseAdapter.Connect()
	if err != nil {
		panic(err)
	}

	authRepository := repository.NewAuthRepository(conn)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signin", authController.SignIn).Methods(http.MethodPost)
	}

	return r
}
