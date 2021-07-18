package private

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func help(b *gotgbot.Bot, c *ext.Context) error {
	c.Message.Reply(
		b,
		"First of all, you need to have your news feed link. "+
			" To get it: visit github.com while you are logged in, "+
			"scroll down until you see the \"Subscribe to your news feed\" link and copy it.\n\n"+
			"After that, send /add_feed followed by your feed link and you're done!\n"+
			"And for removing your feed, simply send /rm_feed.",
		nil,
	)
	return nil
}
