package feed

import (
	"fmt"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/mmcdole/gofeed"

	"github.com/Ryomen-Sukuna/GoGitBot/mongo/feeds"
)

const DELAY = time.Minute * 2

func feedWorker(b *gotgbot.Bot) {
	fp := gofeed.NewParser()

	for {
		feeds2, err := feeds.GetFeeds()

		if err != nil {
			time.Sleep(DELAY)
			continue
		}

		for _, feed := range feeds2 {
			go func(feed feeds.Feed) {
				feed2, err := fp.ParseURL(feed.Url)
				if err != nil {
					return
				}

				if feed2.Items[0].Title == "" || feed2.Items[0].Title == feed.LastTitle {
					return
				}

				feeds.UpdateFeed(feed.ChatId, feed.Url, feed2.Items[0].Title)

				text := fmt.Sprintf(
					"<b>Title</b>: <a href=\"%s\">%s</a>\n",
					feed2.Items[0].Link,
					feed2.Items[0].Title) + fmt.Sprintf(
					"<b>Author</b>: %s", feed2.Items[0].Author.Name,
				)

				if feed2.Items[0].Author.Email != "" {
					text += fmt.Sprintf(
						" &lt;%s&gt;",
						feed2.Items[0].Author.Email,
					)
				}

				text += "\n" + fmt.Sprintf(
					"<b>Published</b>: <code>%s</code>\n", feed2.Items[0].Published) + fmt.Sprintf(
					"<b>Last updated</b>: <code>%s</code>", feed2.Items[0].Updated,
				)

				b.SendMessage(
					feed.ChatId,
					text,
					&gotgbot.SendMessageOpts{
						ParseMode:             "HTML",
						DisableWebPagePreview: true,
					},
				)
			}(feed)
		}

		time.Sleep(DELAY)
	}
}

func StartFeedWorker(b *gotgbot.Bot) {
	go feedWorker(b)
}
