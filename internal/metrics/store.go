package metrics

import "sync"

var (
	mutex           sync.Mutex
	requestCounters = make(map[string]map[string]int)
)

func IncrementRequestCounter(path string, method string) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := requestCounters[path]; !ok {
		requestCounters[path] = make(map[string]int)
	}
	requestCounters[path][method]++
}
