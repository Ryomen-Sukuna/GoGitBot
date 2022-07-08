package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joho/godotenv"

	"github.com/Ryomen-Sukuna/GoGitBot/handlers"
	"github.com/Ryomen-Sukuna/GoGitBot/workers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	b, err := gotgbot.NewBot(
		os.Getenv("TOKEN"),
		&gotgbot.BotOpts{
			Client: http.Client{},
			DefaultRequestOpts: &gotgbot.RequestOpts{
				Timeout: gotgbot.DefaultTimeout,
			},
		},
	)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	workers.StartWorkers(b)
	handlers.AddHandlers(dispatcher)

	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		log.Fatalf("failed to start polling: %v\n", err)
	} else {
		log.Println("Started Polling...!")
	}

	fmt.Printf("%s has been started...!\nMade by @Ryomen-Sukuna\n", b.User.Username)
	updater.Idle()
}
