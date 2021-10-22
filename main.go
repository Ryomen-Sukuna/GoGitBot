package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joho/godotenv"

	"github.com/rojserbest/GitHubFeedsBot/handlers"
	"github.com/rojserbest/GitHubFeedsBot/workers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	b, err := gotgbot.NewBot(
		os.Getenv("TOKEN"),
		&gotgbot.BotOpts{
			Client:      http.Client{},
			GetTimeout:  gotgbot.DefaultGetTimeout,
			PostTimeout: gotgbot.DefaultPostTimeout,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	workers.StartWorkers(b)
	handlers.AddHandlers(dispatcher)

	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		log.Fatal(err)
	}

	updater.Idle()
}
