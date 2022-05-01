package private

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"github.com/Ryomen-Sukuna/GoGitBot/handlers/private/feeds"
)

func AddHandlers(dp *ext.Dispatcher) {
	handlers2 := []ext.Handler{
		handlers.NewCommand("start", start),
		handlers.NewCommand("help", help),
	}

	handlers2 = append(handlers2, feeds.GetHandlers(dp)...)

	dp.AddHandler(
		handlers.NewMessage(
			func(msg *gotgbot.Message) bool {
				return msg.Chat.Type == "private"
			},
			func(b *gotgbot.Bot, c *ext.Context) error {
				for _, handler := range handlers2 {
					if handler.CheckUpdate(b, c.Update) {
						handler.HandleUpdate(b, c)
					}
				}

				return nil
			},
		),
	)
}