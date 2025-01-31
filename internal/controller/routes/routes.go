package routes

import (
	"net/http"
	"os"

	"github.com/BrunoPolaski/login-service/internal/config/logger"
	controllerfactory "github.com/BrunoPolaski/login-service/internal/domain/factory/controller_factory"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	logger.Info("Initializing routes")
	r := mux.NewRouter()
	authController := controllerfactory.GetAuthController()

	r.HandleFunc("/signin", authController.SignIn).Methods(http.MethodPost)

	if os.Getenv("ENV") == "dev" {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}).Methods(http.MethodGet)
	}

	return r
}
