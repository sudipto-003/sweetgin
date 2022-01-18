package repository

import (
	"context"

	"github.com/sudipto-003/sweet-gin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (client *Repos) CreateNewParcelOreder(ctx context.Context, parcel *models.Parcel) error {
	parcelColl := client.MongoClient.Database("sweetgin").Collection("parcel")
	parcel.ID = primitive.NewObjectID()
	res, err := parcelColl.InsertOne(ctx, parcel)
	if err != nil {
		return err
	}

	parcel.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (client *Repos) GetParcelInfoById(ctx context.Context, id primitive.ObjectID, parcel *models.Parcel) error {
	parcelColl := client.MongoClient.Database("sweetgin").Collection("parcel")

	filter := bson.D{{"_id", id}}
	opts := options.FindOne().SetShowRecordID(true)
	if err := parcelColl.FindOne(ctx, filter, opts).Decode(parcel); err != nil {
		return err
	}

	return nil
}

func (client *Repos) GetAllParcels(ctx context.Context) ([]models.Parcel, error) {
	parcelColl := client.MongoClient.Database("sweetgin").Collection("parcel")
	var parcels []models.Parcel

	filter := bson.D{}
	cur, err := parcelColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &parcels); err != nil {
		return nil, err
	}

	return parcels, nil
}
