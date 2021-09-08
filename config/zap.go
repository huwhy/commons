package config

type Zap struct {
	Level       string `yaml:"level"`
	Dir         string `yaml:"dir"`
	ShowLine    bool   `yaml:"show_line"`
	EncodeLevel string `yaml:"encode_level"`
	ShowConsole bool   `yaml:"show_console"`
	Format      string `yaml:"format"`
}
