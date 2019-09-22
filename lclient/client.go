package mclient

import (
	"sync"
)

var (
	clients map[string]interface{}
	lock    sync.RWMutex
)

// Init Init client
func Init() {
	clients = map[string]interface{}{}
}

// Add Add client
func Add(k string, c interface{}) {
	lock.Lock()
	defer lock.Unlock()

	clients[k] = c
}

// Get Get client
func Get(key string) interface{} {
	lock.Lock()
	defer lock.Unlock()

	if _, ok := clients[key]; ok {
		return clients[key]
	}
	return nil
}

// All All clients
func All() map[string]interface{} {
	return clients
}

// Del Del client
func Del(k string) {
	lock.Lock()
	defer lock.Unlock()

	delete(clients, k)
}
