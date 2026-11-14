package api_runner

import (
	"sync"
	"time"
)

const requestInterval = 250 * time.Millisecond

var (
	requestMu   sync.Mutex
	lastRequest time.Time
)

// WaitForRequest ограничивает частоту запросов к учебному стенду без повторов запросов.
func WaitForRequest() {
	requestMu.Lock()
	defer requestMu.Unlock()
	if delay := time.Until(lastRequest.Add(requestInterval)); delay > 0 {
		time.Sleep(delay)
	}
	lastRequest = time.Now()
}
