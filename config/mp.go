package config

type MpConfig struct {
	AppId  string `yaml:"appId"`
	Secret string `yaml:"secret"`
	Token  string `yaml:"token"`
	AesKey string `yaml:"aesKey"`
}
