package feeds

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	mongo2 "github.com/rojserbest/GitHubFeedsBot/mongo"
)

const MAX_FEEDS = 5

var collection *mongo.Collection

type Feed struct {
	ChatId    int64  `bson:"chatId"`
	Url       string `bson:"url"`
	LastTitle string `bson:"lastTitle"`
}

func getCollection() *mongo.Collection {
	if collection != nil {
		return collection
	}

	collection = mongo2.GetDatabase().Collection("feeds")
	return collection
}

func AddFeed(feed Feed) error {
	_, err := getCollection().InsertOne(mongo2.Ctx, feed)
	return err
}

func UpdateFeed(chatId int64, url string, lastTitle string) error {
	_, err := getCollection().UpdateOne(mongo2.Ctx, bson.M{"chatId": chatId, "url": url}, bson.M{"$set": bson.M{"lastTitle": lastTitle}})
	return err
}

func DeleteFeed(chatId int64, url string) bool {
	del, _ := getCollection().DeleteOne(mongo2.Ctx, bson.M{"chatId": chatId, "url": url})
	return del.DeletedCount != 0
}

func DeleteFeeds(chatId int64) int64 {
	del, _ := getCollection().DeleteOne(mongo2.Ctx, bson.M{"chatId": chatId})
	return del.DeletedCount
}

func GetFeeds() ([]Feed, error) {
	feeds := []Feed{}
	find, err := getCollection().Find(mongo2.Ctx, bson.M{"chatId": bson.M{"$exists": 1}})
	if err != nil {
		return feeds, err
	}

	return feeds, find.All(mongo2.Ctx, &feeds)
}

func GetFeedsForChat(chatId int64) ([]Feed, error) {
	feeds := []Feed{}
	find, err := getCollection().Find(mongo2.Ctx, bson.M{"chatId": chatId})
	if err != nil {
		return feeds, err
	}

	return feeds, find.All(mongo2.Ctx, &feeds)
}

func HasFeed(chatId int64, url string) bool {
	url = strings.Replace(url, "http://", "", 1)
	url = strings.Replace(url, "https://", "", 1)

	feeds, _ := GetFeedsForChat(chatId)

	for _, feed := range feeds {
		feed.Url = strings.Replace(feed.Url, "http://", "", 1)
		feed.Url = strings.Replace(feed.Url, "https://", "", 1)

		if feed.Url == url {
			return true
		}
	}

	return false
}

func GetFeedsCount() int {
	feeds, _ := GetFeeds()
	return len(feeds)
}

func GetFeedsCountForChat(chatId int64) int {
	feeds, _ := GetFeedsForChat(chatId)
	return len(feeds)
}

func HasEnoughFeeds(chatId int64) bool {
	return GetFeedsCountForChat(chatId) > MAX_FEEDS
}
