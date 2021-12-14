package workers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"

	"github.com/roj1512/GitHubFeedBot/workers/feed"
)

func StartWorkers(b *gotgbot.Bot) {
	feed.StartFeedWorker(b)
}
