package routes

import (
	"net/http"

	controllerfactory "github.com/BrunoPolaski/login-service/internal/domain/factory/controller_factory"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()
	authController := controllerfactory.GetAuthController()

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signin", authController.SignIn).Methods(http.MethodPost)

	return r
}
