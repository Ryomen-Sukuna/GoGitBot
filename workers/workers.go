package workers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"

	"FeedsBot/workers/feed"
)

func StartWorkers(b *gotgbot.Bot) {
	feed.StartFeedWorker(b)
}
