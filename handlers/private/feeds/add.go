package feeds

import (
	"net/http"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"FeedsBot/mongo/feeds"
)

func add(b *gotgbot.Bot, c *ext.Context) error {
	if feeds.HasEnoughFeeds(c.EffectiveChat.Id) {
		c.Message.Reply(b, "You can't have more feeds.", nil)
		return nil
	}

	fields := strings.Fields(c.Message.Text)
	if len(fields) != 2 {
		c.Message.Reply(b, "Invalid arguments provided.", nil)
		return nil
	}

	url := fields[1]
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		c.Message.Reply(b, "Invalid URL provided.", nil)
		return nil
	}

	if feeds.HasFeed(c.EffectiveChat.Id, url) {
		c.Message.Reply(b, "You already have this feed.", nil)
		return nil
	}

	err = feeds.AddFeed(feeds.Feed{ChatId: c.EffectiveChat.Id, Url: url})
	if err != nil {
		c.Message.Reply(b, "An error occurred adding the feed: "+err.Error(), nil)
		return nil
	}

	c.Message.Reply(b, "Feed added.", nil)
	return nil
}
