package core

import (
	"github.com/huwhy/commons/config"
	"testing"
)

func TestLoadConf(t *testing.T) {
	LoadConfig()
}

func TestLoadConfig(t *testing.T) {
	conf := new(KeepConfig)
	LoadConf(conf)
	t.Log(conf)
}

type KeepConfig struct {
	Server  *config.Server   `yaml:"server"`
	Mysql   *config.Mysql    `yaml:"mysql"`
	Zap     *config.Zap      `yaml:"zap"`
	JWT     *config.JWT      `yaml:"jwt"`
	Mp      *config.MpConfig `yaml:"mp"`
	ROOT    string
	Profile string
}

func (c *KeepConfig) SetRoot(root string) {
	c.ROOT = root
}

func (c *KeepConfig) SetEnv(env string) {
	c.Profile = env
}

func (c *KeepConfig) SetPort(port string) {
	c.Server.Port = port
}

func (c *KeepConfig) IsDev() bool {
	return c.Profile == "dev"
}
