package lcfg

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcfg"
)

// Tree Tree
var (
	// Name Name
	Name string
	// Tree Tree
	Tree *gcfg.Config
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
func Instance(name ...string) *gcfg.Config {
	return g.Config(name...)
}

// InitCfg config init
func InitCfg() (cfg *gcfg.Config, err error) {
	if Name != "" {
		g.Config().SetFileName(Name)
	}

	cfg = g.Config()
	if cfg.FilePath() == "" || cfg.GetFileName() == "" {
		return nil, fmt.Errorf("配置文件( %s )不存在", cfg.GetFileName())
	}

	Tree = cfg
	return
}

// GetCfg get config
func GetCfg() *gcfg.Config {
	return Tree
}
