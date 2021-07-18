package feeds

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"FeedsBot/mongo/feeds"
)

func rm(b *gotgbot.Bot, c *ext.Context) error {
	del := feeds.DeleteFeed(c.EffectiveChat.Id)
	if del {
		c.Message.Reply(b, "Feed deleted.", nil)
		return nil
	}

	c.Message.Reply(b, "Feed not deleted.", nil)
	return nil
}
