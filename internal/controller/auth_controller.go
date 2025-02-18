package controller

import (
	"net/http"

	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/BrunoPolaski/login-service/internal/domain/service"
)

type AuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &authController{}
}

func (ac authController) SignIn(w http.ResponseWriter, r *http.Request) {
	logger.Info("Authenticating user")
	_, _, ok := r.BasicAuth()
	if !ok {
		logger.Error("Basic auth header not found")
		http.Error(w, "Basic auth header not found", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
