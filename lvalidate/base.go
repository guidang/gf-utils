package mvalidate

// Mode 生成校验码的方式
type Mode int

const (
	// ByteOverlayAccumulation 字节加密方式
	ByteOverlayAccumulation Mode = iota
)
