package metrics

import "sync"

var (
	mutex           sync.Mutex
	requestCounters = make(map[string]int)
)

func IncrementRequestCounter(path string) {
	mutex.Lock()
	defer mutex.Unlock()
	requestCounters[path]++
}
