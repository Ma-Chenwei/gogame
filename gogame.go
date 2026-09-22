package gogame

import "sync"

var (
	initialized bool
	running     bool
	mu          sync.RWMutex
)

func Init() {
	mu.Lock()
	defer mu.Unlock()

	if initialized {
		return
	}

	initialized = true
	running = true

	initEvents()
}

func Quit() {
	mu.Lock()
	running = false
	initialized = false
	mu.Unlock()
}

func Running() bool {
	mu.RLock()
	defer mu.RUnlock()

	return running
}

func GetInit() bool {
	mu.RLock()
	defer mu.RUnlock()

	return initialized
}