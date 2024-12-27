package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/BrunoPolaski/login-service/internal/config/logger"
)

type AuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request)
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (ac authController) SignIn(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	username, password, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(map[string]string{
			"error": "unauthorized",
		})
		return
	}

	if username != os.Getenv("USERNAME") || password != os.Getenv("PASSWORD") {
		logger.Warn(
			fmt.Sprintf("Unauthorized access attempt from %s", r.RemoteAddr),
		)

		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(map[string]string{
			"error": "unauthorized",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}