package private

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func start(b *gotgbot.Bot, c *ext.Context) error {
	c.Message.Reply(
		b,
		"I can send your GitHub activities directly to here. Use /help to know how to use me.",
		nil,
	)
	return nil
}
