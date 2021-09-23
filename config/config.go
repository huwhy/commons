package config

type IConfig interface {
	SetRoot(root string)
	SetEnv(env string)
	SetPort(port string)
	IsDev() bool
}

type Config struct {
	Server  *Server   `yaml:"server"`
	Mysql   *Mysql    `yaml:"mysql"`
	Zap     *Zap      `yaml:"zap"`
	JWT     *JWT      `yaml:"jwt"`
	Mp      *MpConfig `yaml:"mp"`
	ROOT    string
	Profile string
}

func (c Config) SetRoot(root string) {
	c.ROOT = root
}

func (c Config) SetEnv(env string) {
	c.Profile = env
}

func (c Config) SetPort(port string) {
	c.Server.Port = port
}

func (c Config) IsDev() bool {
	return c.Profile == "dev"
}
