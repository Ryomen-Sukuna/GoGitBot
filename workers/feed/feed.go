package feed

import (
	"fmt"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
	"strconv"
)

const DELAY = time.Minute * 2

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

			text := fmt.Sprintf("<b>Title</b>: <a href=\"%s\">%s</a>\n", feed.Items[0].Link, feed.Items[0].Title) +
				fmt.Sprintf("<b>Author</b>: %s", feed.Items[0].Authors[0].Name)

			if feed.Items[0].Authors[0].Email != "" {
				text += fmt.Sprintf(" &lt;%s&gt;", feed.Items[0].Authors[0].Name)
			}

			text += "\n" + fmt.Sprintf("<b>Published</b>: <code>%s</code>\n", feed.Items[0].Published) +
				fmt.Sprintf("<b>Last updated</b>: <code>%s</code>", feed.Items[0].Updated)

			b.SendMessage(chatId, text, &gotgbot.SendMessageOpts{ParseMode: "HTML", DisableWebPagePreview: true})
		}()

		time.Sleep(DELAY)
	}
}

func StartFeedWorker(b *gotgbot.Bot) {
	go feedWorker(b)
}
