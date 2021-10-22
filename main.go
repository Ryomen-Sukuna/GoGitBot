package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joho/godotenv"

	"github.com/rojserbest/GitHubFeedBot/workers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	b, err := gotgbot.NewBot(
		os.Getenv("BOT_TOKEN"),
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

	workers.StartWorkers(b)

	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		log.Fatal(err)
	}

	updater.Idle()
}
