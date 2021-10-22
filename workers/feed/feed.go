package feed

import (
	"fmt"
	"github.com/xeonx/timeago"
	"html"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
	"strconv"
)

const DELAY = time.Minute * 2

func parseTime(t string) string {
	parsed, err := time.Parse(time.RFC3339, t)

	if err != nil {
		return t
	}

	return timeago.English.Format(parsed)
}

func feedWorker(b *gotgbot.Bot) {
	fp := gofeed.NewParser()

	chatIdString := os.Getenv("CHAT_ID")
	chatIdInt, err := strconv.Atoi(chatIdString)
	if err != nil {
		log.Fatal(err)
	}

	chatId := int64(chatIdInt)
	url := os.Getenv("FEED_URL")
	lastTitle := ""

	for {
		if err != nil {
			time.Sleep(DELAY)
			continue
		}

		go func() {
			feed, err := fp.ParseURL(url)
			if err != nil {
				return
			}

			if feed.Items[0].Title == "" || feed.Items[0].Title == lastTitle {
				return
			}

			lastTitle = feed.Items[0].Title

			text := fmt.Sprintf("<b>Title</b>: <a href=\"%s\">%s</a>\n", feed.Items[0].Link, html.EscapeString(feed.Items[0].Title)) +
				fmt.Sprintf("<b>Author</b>: <a href=\"https://github.com/%s\">%s</a>", feed.Items[0].Authors[0], html.EscapeString(feed.Items[0].Authors[0].Name))

			if feed.Items[0].Authors[0].Email != "" {
				text += fmt.Sprintf(" &lt;%s&gt;", html.EscapeString(feed.Items[0].Authors[0].Email))
			}

			text += "\n" + fmt.Sprintf(parseTime(feed.Items[0].Published))

			b.SendMessage(chatId, text, &gotgbot.SendMessageOpts{ParseMode: "HTML", DisableWebPagePreview: true})
		}()

		time.Sleep(DELAY)
	}
}

func StartFeedWorker(b *gotgbot.Bot) {
	go feedWorker(b)
}
