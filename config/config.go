package config

import (
	"fmt"
	"os"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	if err := c.initConfig(); err != nil {
		return err
	}
	c.watchConfig()
	return nil
}

func (c *Config) initConfig() error{
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		goenv := os.Getenv("GO_ENV")
		if goenv == "" {
			goenv = "development"
		}
		viper.AddConfigPath("config")
        viper.SetConfigName(goenv)
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
        fmt.Printf("Config file changed: %s\n", e.Name)
    })
}