package ldb

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// Init Init DB
func Init() error {
	return Ping()
}

// Get Get DB
func Get() (db gdb.DB) {
	return g.DB()
}

// Ping 检测数据库连接是否正常
func Ping() error {
	if err := Get().PingMaster(); err != nil {
		return fmt.Errorf("数据库连接失败: %s", err.Error())
	}
	return nil
}
