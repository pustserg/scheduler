package tgbot

import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/pustserg/scheduler/appconfig"
	"log"
	"sync"
)

var botInstance *Tgbot
var once sync.Once

// Tgbot wrapper around telegram-bot-api
type Tgbot struct {
	api     *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
	Status  string
}

// Init initialize bot
func (telegramBot *Tgbot) Init(env string) {
	cfg := appconfig.LoadConfig(env)
	log.Println("cfg received", cfg.TelegramAPIKey)
	botAPI, err := tgbotapi.NewBotAPI(cfg.TelegramAPIKey)
	if err != nil {
		panic(err)
	}
	telegramBot.api = botAPI
	log.Println("bot api initiated")
	botUpdate := tgbotapi.NewUpdate(0)
	botUpdate.Timeout = 64
	botUpdates, err := telegramBot.api.GetUpdatesChan(botUpdate)
	if err != nil {
		log.Fatal(err)
	}
	telegramBot.Updates = botUpdates
	telegramBot.Status = "initiated"
}
