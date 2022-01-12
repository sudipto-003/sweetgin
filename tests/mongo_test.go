package tests

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoConnection(t *testing.T) {
	const uri = "mongodb://0.0.0.0:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatalf("%q\n", err)
	}
	defer client.Disconnect(context.TODO())

	if err := client.Ping(context.TODO(), nil); err != nil {
		t.Fatalf("Ping Unsuccessful. Error %q\n", err)
	}
}
