package ldb

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// InitDB InitDB
func InitDB() (err error) {
	err = Ping()
	return
}

// GetDB GetDB
func GetDB() (db gdb.DB) {
	db = g.DB()
	return db
}

// Ping 检测数据库连接是否正常
func Ping() (err error) {
	if err = GetDB().PingMaster(); err != nil {
		return fmt.Errorf("数据库连接失败: %s", err.Error())
	}
	return err
}
