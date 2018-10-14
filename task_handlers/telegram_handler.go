package taskHandlers

import (
	"github.com/pustserg/scheduler/appconfig"
	"github.com/pustserg/scheduler/tasks"
	"github.com/pustserg/scheduler/tgbot"
	"log"
)

// Handle runs handler for task
func Handle(task *tasks.Task, cfg *appconfig.Config) {
	switch task.Action {
	case "send_telegram_message":
		handleTelegramMessageTask(task)
	default:
		log.Println("I cant handle task with action", task.Action)
	}
}

func handleTelegramMessageTask(task *tasks.Task) {
	bot := tgbot.Tgbot{}
	bot.Send("Hello")
}
