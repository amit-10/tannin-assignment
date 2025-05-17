package metrics

import (
	"fmt"
	"net/http"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Content-Type", "text/plain")

	for requestKey, count := range requestCounters {
		fmt.Fprintf(w, "petstore_http_requests_total{path=\"%s\", method=\"%s\"} %d\n", requestKey.Path, requestKey.Method, count)
	}
}
