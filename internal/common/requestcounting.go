package common

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tanninio/home-assignment/internal/metrics"
)

func RequestCountingMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			route := mux.CurrentRoute(r)
			path, err := route.GetPathTemplate()
			if err == nil && path != "/metrics" {
				metrics.IncrementRequestCounter(path)
			}
			next.ServeHTTP(w, r)
		})
	}
}
