package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
)

func BasicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		restErr := rest_err.NewBadRequestError("authentication is invalid")

		auth := r.Header.Get("Authorization")
		if !strings.Contains(auth, "Basic ") {
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)

			return
		}

		if len(strings.Split(auth, " ")) != 2 {
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)

			return
		}

		next.ServeHTTP(w, r)
	})
}
