package lcfg

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcfg"
)

var (
	// Name Name
	Name string
)

// SetCfgName SetCfgName
func SetCfgName(n string) {
	Name = n
}

// GetCfgName GetCfgName
func GetCfgName() string {
	return Name
}

// Instance Instance

// InitCfg config init
func InitCfg() (*gcfg.Config, error) {
	if Name != "" {
		g.Config().SetFileName(Name)
	}

	cfg := g.Config()
	if cfg.FilePath() == "" || cfg.GetFileName() == "" {
		return nil, fmt.Errorf("配置文件( %s )不存在", cfg.GetFileName())
	}

	return cfg, nil
}

// GetCfg get config
func GetCfg(name ...string) *gcfg.Config {
	return g.Config(name...)
}
