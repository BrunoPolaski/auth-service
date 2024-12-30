package routes

import (
	"net/http"
	"os"

	controllerfactory "github.com/BrunoPolaski/login-service/internal/domain/factory/controller_factory"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()
	authController := controllerfactory.GetAuthController()

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signin", authController.SignIn).Methods(http.MethodPost)

	if os.Getenv("ENV") == "dev" {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}).Methods(http.MethodGet)
	}

	return r
}
