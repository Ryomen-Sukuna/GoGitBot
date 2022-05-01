package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	Ctx    context.Context
	err    error
)

func GetClient() *mongo.Client {
	if Client != nil {
		return Client
	}

	Client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO")))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Client.Connect(Ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return Client
}

func GetDatabase() *mongo.Database {
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "feedsbot"
	}

	return GetClient().Database(name)
}
