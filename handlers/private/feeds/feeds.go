package feeds

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func GetHandlers(dp *ext.Dispatcher) []ext.Handler {
	return []ext.Handler{
		handlers.NewCommand("rm_feed", rm),
		handlers.NewCommand("rm_feeds", rmAll),
		handlers.NewCommand("add_feed", add),
	}
}
