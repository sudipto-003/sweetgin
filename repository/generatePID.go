package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (client *Repos) GetNewPID() string {
	pidcoll := client.MongoClient.Database("sweetgin").Collection("pids")
	prefix := time.Now().Format("060102")
	index := bson.M{}

	filter := bson.D{{"_id", prefix}}
	idx_inc := bson.D{
		{"$inc",
			bson.D{
				{"index", 1},
			},
		},
	}
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)
	err := pidcoll.FindOneAndUpdate(context.Background(), filter, idx_inc, opts).Decode(&index)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			pidcoll.InsertOne(context.Background(), bson.D{{"_id", prefix}, {"index", 1}})
			index["index"] = 1
		}
	}

	return prefix + fmt.Sprintf("%04d", index["index"])
}
