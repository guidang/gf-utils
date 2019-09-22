package mclient

import (
	"github.com/gogf/gf/container/gmap"
	"sync"
)

var (
	clients gmap.StrAnyMap
	lock    sync.RWMutex
)

// Init Init client
func Init() {
	clients = gmap.StrAnyMap{}
}

// Add Add client
func Add(k string, c interface{}) {
	lock.Lock()
	defer lock.Unlock()

	clients.Set(k, c)
}

// Get Get client
func Get(k string) interface{} {
	lock.Lock()
	defer lock.Unlock()

	return clients.Get(k)
}

// All All clients
func All() gmap.StrAnyMap {
	return clients
}

// Del Del client
func Del(k string) {
	lock.Lock()
	defer lock.Unlock()

	clients.Remove(k)
}
