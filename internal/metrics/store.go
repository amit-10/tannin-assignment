package metrics

import "sync"

type RouteKey struct {
	Path   string
	Method string
}

var (
	mutex           sync.Mutex
	requestCounters = make(map[RouteKey]int)
)

func IncrementRequestCounter(path string, method string) {
	mutex.Lock()
	defer mutex.Unlock()
	routeKey := RouteKey{Path: path, Method: method}
	requestCounters[routeKey]++
}
