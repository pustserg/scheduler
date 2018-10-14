package tgbot

import (
	"context"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/pustserg/scheduler/appconfig"
	"golang.org/x/net/proxy"
	"log"
	"net"
	"net/http"
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
	dialer, proxyErr := proxy.SOCKS5(
		"tcp",
		cfg.TelegramProxyURL,
		&proxy.Auth{User: cfg.TelegramProxyUser, Password: cfg.TelegramProxyPassword},
		proxy.Direct,
	)
	if proxyErr != nil {
		log.Panicf("Error in proxy %s", proxyErr)
	}
	client := &http.Client{
		Transport: &http.Transport{DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
		},
	}
	botAPI, err := tgbotapi.NewBotAPIWithClient(cfg.TelegramAPIKey, client)
	if err != nil {
		log.Panicf("Error in bot %s", err)
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
	for update := range telegramBot.Updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		botAPI.Send(msg)
	}
}

// Send sends a message
func (telegramBot *Tgbot) Send(message string) {
	telegramBot.Send(message)
}
