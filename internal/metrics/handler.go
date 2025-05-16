package metrics

import (
	"fmt"
	"net/http"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Content-Type", "text/plain")

	for path, count := range requestCounters {
		fmt.Fprintf(w, "petstore_http_requests_total{path=\"%s\"} %d\n", path, count)
	}
}
