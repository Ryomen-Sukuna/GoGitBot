package workers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"

	"github.com/rojserbest/GitHubFeedBot/workers/feed"
)

func StartWorkers(b *gotgbot.Bot) {
	feed.StartFeedWorker(b)
}
