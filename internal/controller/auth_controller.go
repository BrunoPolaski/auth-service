package controller

import (
	"encoding/json"
	"fmt"
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
	return &authController{
		authService: service,
	}
}

func (ac *authController) SignIn(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	logger.Info("Authenticating user")
	username, password, ok := r.BasicAuth()
	if !ok {
		logger.Error("Basic auth header not found")
		http.Error(w, "Basic auth header not found", http.StatusUnauthorized)
		return
	}

	token, err := ac.authService.SignIn(username, password)
	if err != nil {
		logger.Error(fmt.Sprintf("Error signing in user: %s", err.Error()))
		encoder.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	encoder.Encode(map[string]string{
		"token": token,
	})

	logger.Info("User authenticated")
}
