package lclient

import (
	"github.com/gogf/gf/container/gmap"
)

var (
	clients *gmap.StrAnyMap
)

// Init Init client
func Init() {
	clients = gmap.NewStrAnyMap(true)
}

// Add Add client
func Add(k string, c interface{}) {
	clients.Set(k, c)
}

// Get Get client
func Get(k string) interface{} {
	return clients.Get(k)
}

// All All clients
func All() *gmap.StrAnyMap {
	return clients
}

// Del Del client
func Del(k string) {
	clients.Remove(k)
}
