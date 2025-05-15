package common

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	mutex           sync.RWMutex
	requestCounters = map[string]int{}
)

func RequestCountingMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			route := mux.CurrentRoute(r)
			path, err := route.GetPathTemplate()
			if err == nil {
				mutex.Lock()
				requestCounters[path]++
				mutex.Unlock()
			}
			next.ServeHTTP(w, r)
		})
	}
}
