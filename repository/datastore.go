package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repos struct {
	MongoClient *mongo.Client
	RedisClient *redis.Client
}

func CreateMongoConnection(ctx context.Context, uri string) (*mongo.Client, error) {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return conn, err
}

func CreateRedisConnection(ctx context.Context, url string) (*redis.Client, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return redis.NewClient(opts), nil
}
