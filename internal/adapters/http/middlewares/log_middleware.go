package middlewares

import (
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/auth-service/internal/infra/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(
			fmt.Sprintf(
				"----------------- Received request ----------------- \nMethod: %s\nPath: %s\nHeaders: %v\n",
				r.Method,
				r.URL.Path,
				r.Header,
			),
		)
		next.ServeHTTP(w, r)
	})
}
