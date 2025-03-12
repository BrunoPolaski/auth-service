package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/BrunoPolaski/login-service/internal/thirdparty/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Info(
				fmt.Sprintf("%s \t %s \t %s", r.Method, r.URL.Path, time.Since(start)),
			)
		},
	)
}
