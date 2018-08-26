package appconfig

import (
	"github.com/jinzhu/configor"
)

// Config fof app
type Config struct {
	DbFile string
}

//LoadConfig returns config
func LoadConfig(env string) *Config {
	config := Config{}
	fileName := "config." + env + ".yml"
	configor.Load(&config, fileName)
	return &config
}
