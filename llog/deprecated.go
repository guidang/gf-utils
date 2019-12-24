package llog

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
