package routes

import (
	"net/http"

	"github.com/BrunoPolaski/login-service/internal/config/database"
	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/BrunoPolaski/login-service/internal/domain/factory"
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

	controllerFactory := factory.NewControllerFactory(conn)

	authController := controllerFactory.GetAuthController()
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signin", authController.SignIn).Methods(http.MethodPost)
	}

	return r
}
