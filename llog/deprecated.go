package llog

import (
	"github.com/gogf/gf/os/glog"
)

// GetLog GetLog
// Deprecated
func GetLog() *glog.Logger {
	return Log
}

// InitLog log init
// Deprecated
func InitLog() error {
	return Init()
}
