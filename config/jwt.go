package config

type JWT struct {
	Sign    string `yaml:"sign"`
	Expires int64  `yaml:"expires"`
}
