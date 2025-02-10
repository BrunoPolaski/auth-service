package routes

import (
	"net/http"

	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	logger.Info("Initializing routes")
	r := mux.NewRouter()
	authController := controllerfactory.GetAuthController()

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signin", authController.SignIn).Methods(http.MethodPost)
	}

	return r
}
