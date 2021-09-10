package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/huwhy/commons/config"
	"github.com/spf13/viper"
	"path/filepath"
)

func LoadConfig() (*viper.Viper, *config.Config) {
	root, _ := filepath.Abs("./")
	var configPath, profile, port string
	flag.StringVar(&profile, "env", "dev", "启动环境")
	flag.StringVar(&port, "port", "", "端口")
	flag.StringVar(&configPath, "c", "", "choose config file.")
	flag.Parse()
	if configPath == "" {
		configPath = root + "/config.yml"
	}
	var conf = new(config.Config)
	conf.ROOT = root
	conf.Profile = profile
	v := viper.New()
	v.SetConfigFile(configPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(conf); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(conf); err != nil {
		fmt.Println(err)
	}
	if port != "" {
		conf.Server.Port = port
	}
	return v, conf
}

func LoadConf(conf config.IConfig) *viper.Viper {
	root, _ := filepath.Abs("./")
	var configPath, profile, port string
	flag.StringVar(&profile, "env", "dev", "启动环境")
	flag.StringVar(&port, "port", "", "端口")
	flag.StringVar(&configPath, "c", "", "choose config file.")
	flag.Parse()
	if configPath == "" {
		configPath = root + "/config.yml"
	}
	conf.SetRoot(root)
	conf.SetEnv(profile)
	v := viper.New()
	v.SetConfigFile(configPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(conf); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&conf); err != nil {
		fmt.Println(err)
	}
	if port != "" {
		conf.SetPort(port)
	}
	return v
}
