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

func init() {
	c := Config{
		Name: "",
	}
	if err := c.initConfig(); err != nil {
		panic(err)
	}
	c.watchConfig()
}

func (c *Config) initConfig() error {
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
