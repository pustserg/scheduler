package appconfig

import (
	"github.com/jinzhu/configor"
	"log"
	"sync"
)

// Config fof app
type Config struct {
	DbFile                string
	TelegramAPIKey        string
	TelegramChatID        int64
	TelegramProxyURL      string
	TelegramProxyUser     string
	TelegramProxyPassword string
}

var cfgInstance *Config
var once sync.Once

//LoadConfig returns config
func LoadConfig(env string) *Config {
	once.Do(func() {
		log.Println("Cfg init once")
		config := Config{}
		configor.New(&configor.Config{Environment: env}).Load(&config, "config.yml")
		cfgInstance = &config
	})
	return cfgInstance
}
