package feeds

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	mongo2 "FeedsBot/mongo"
)

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

func UpdateFeed(chatId int64, lastTitle string) error {
	_, err := getCollection().UpdateOne(mongo2.Ctx, bson.M{"chatId": chatId}, bson.M{"$set": bson.M{"lastTitle": lastTitle}})
	return err
}

func DeleteFeed(chatId int64) bool {
	del, _ := getCollection().DeleteOne(mongo2.Ctx, bson.M{"chatId": chatId})
	return del.DeletedCount != 0
}

func GetFeeds() ([]Feed, error) {
	feeds := []Feed{}
	find, err := getCollection().Find(mongo2.Ctx, bson.M{"chatId": bson.M{"$exists": 1}})
	if err != nil {
		return feeds, err
	}

	return feeds, find.All(mongo2.Ctx, &feeds)
}

func GetFeedsCount() int {
	feeds, _ := GetFeeds()
	return len(feeds)
}

func HasFeed(chatId int64) bool {
	feed := Feed{}
	getCollection().FindOne(mongo2.Ctx, bson.M{"chatId": chatId}).Decode(&feed)
	return feed != Feed{}
}
