package main

import (
	"flag"
	"fmt"
	"github.com/pustserg/scheduler/appconfig"
	"github.com/pustserg/scheduler/daemon"
	"github.com/pustserg/scheduler/server"
	"github.com/pustserg/scheduler/tasks"
	"github.com/pustserg/scheduler/tgbot"
	"log"
)

func main() {
	daemonWorkers := flag.Int("daemon-workers", 1, "Integer count of daemon workers")
	appenv := string(*flag.String("env", "dev", "App environment"))

	flag.Parse()
	config := appconfig.LoadConfig(appenv)
	log.Println("app started with app env", appenv)
	repo := tasks.NewRepository(config.DbFile)
	startDaemon(*daemonWorkers, repo)
	startHTTPServer(appenv, repo)
	bot := startTgBot(appenv)
	log.Println("bot status", bot.Status)

	var input string
	fmt.Scanln(&input)
}

func startDaemon(workersCount int, repo *tasks.TaskRepository) {
	daemonInstance := daemon.Daemon{}
	fmt.Println(daemonInstance.State)
	daemonInstance.Start(workersCount, repo)
}

func startHTTPServer(env string, repo *tasks.TaskRepository) {
	log.Println("start http server")
	serverInstance := server.NewServer(env, repo)
	log.Println("Http server started with status", serverInstance.Status)
	go serverInstance.Start()
}

func startTgBot(appenv string) *tgbot.Tgbot {
	bot := tgbot.Tgbot{}
	bot.Init(appenv)
	return &bot
}
