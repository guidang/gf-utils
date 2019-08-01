package ucfg

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/gcfg"
)

var cfg *gcfg.Config

// InitCfg config init
func InitCfg() *gcfg.Config {
	cfg = g.Config()
	return cfg
}

// GetCfg get config
func GetCfg() *gcfg.Config {
	return cfg
}
