package feeds

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"FeedsBot/mongo/feeds"
)

func rm(b *gotgbot.Bot, c *ext.Context) error {
	args := strings.Fields(c.Message.Text)
	if len(args) != 2 {
		c.Message.Reply(b, "Invalid arguments or URL not provided.", nil)
		return nil
	}

	del := feeds.DeleteFeed(c.EffectiveChat.Id, args[1])
	if del {
		c.Message.Reply(b, "Feed deleted.", nil)
		return nil
	}

	c.Message.Reply(b, "Feed not deleted.", nil)
	return nil
}

func rmAll(b *gotgbot.Bot, c *ext.Context) error {
	c.Message.Reply(
		b,
		fmt.Sprintf("Deleted %d feeds.", feeds.DeleteFeeds(c.EffectiveChat.Id)),
		nil,
	)
	return nil
}
