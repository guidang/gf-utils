package lredis

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/g"
)

// Init 初始化 Redis
func Init() error {
	return Ping()
}

// Get 获取 Redis
func Get() *gredis.Redis {
	return g.Redis()
}

// Ping Ping
func Ping() error {
	return nil
}
