package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/joho/godotenv"

	"github.com/Ryomen-Sukuna/GoGitBot/workers"
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
		panic("failed to create new bot: " + err.Error())
	}

	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	dispatcher.AddHandler(handlers.NewCommand("start", start))

	workers.StartWorkers(b)

	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		log.Fatalf("failed to start polling: %v\n", err)
	} else {
		log.Println("Started Polling...!")
	}

	fmt.Printf("%s has been started...!\nMade by @Ryomen-Sukuna\n", b.User.Username)
	updater.Idle()
}

func start(bot *gotgbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	user_name := ctx.EffectiveUser.FirstName

	if ctx.EffectiveChat.Type != "private" {
		return ext.EndGroups
	}

	msg.Reply(
		bot,
		fmt.Sprintf(
			"*Hi %v*,\n"+
				"I am a simple bot that notifies you about GitHub activities on Telegram using [Go](https://go.dev)*",
			user_name,
		),
		&gotgbot.SendMessageOpts{
			ParseMode:                "Markdown",
			ReplyToMessageId:         msg.MessageId,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		},
	)
	return ext.EndGroups
}
