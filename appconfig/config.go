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
	configor.New(&configor.Config{Environment: env}).Load(&config, "config.yml")
	return &config
}
