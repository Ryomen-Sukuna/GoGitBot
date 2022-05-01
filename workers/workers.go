package workers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"

	"github.com/Ryomen-Sukuna/GoGitBot/workers/feed"
)

func StartWorkers(b *gotgbot.Bot) {
	feed.StartFeedWorker(b)
}
