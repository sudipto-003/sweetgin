package repository

import (
	"context"
	"time"

	"github.com/sudipto-003/sweet-gin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (client *Repos) CreateNewParcelOreder(ctx context.Context, parcel *models.Parcel) error {
	parcelColl := client.MongoClient.Database("sweetgin").Collection("parcel")
	parcel.ID = primitive.NewObjectID()
	parcel.ParcelId = client.GetNewPID()
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
	if err := parcelColl.FindOne(ctx, filter).Decode(parcel); err != nil {
		return err
	}

	return nil
}

func (client *Repos) GetAllParcels(ctx context.Context, parcels *[]models.Parcel) error {
	parcelColl := client.MongoClient.Database("sweetgin").Collection("parcel")

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"date_picked", 1}})
	cur, err := parcelColl.Find(ctx, filter, opts)
	if err != nil {
		return err
	}

	err = cur.All(ctx, parcels)
	return err
}

func (client *Repos) GetParcelByPID(ctx context.Context, pid string, parcel *models.Parcel) error {
	parcelColl := client.MongoClient.Database("sweetgin").Collection("parcel")

	filter := bson.D{{"parcelid", pid}}
	if err := parcelColl.FindOne(ctx, filter).Decode(parcel); err != nil {
		return err
	}

	return nil
}

func (client *Repos) GetParcelByDate(ctx context.Context, date time.Time, parcels *[]models.Parcel) error {
	parcelColl := client.MongoClient.Database("sweetgin").Collection("parcel")
	filter := bson.D{
		{"date_picked",
			bson.D{
				{"$gte", date.UnixMilli()},
				{"$lt", date.AddDate(0, 0, 1).UnixMilli()},
			},
		},
	}

	cur, err := parcelColl.Find(ctx, filter)
	if err != nil {
		return err
	}

	err = cur.All(ctx, parcels)
	return err
}

func (client *Repos) UpdateParcelStatus(ctx context.Context, pid string, parcel *models.Parcel) error {
	parcelColl := client.MongoClient.Database("sweetgin").Collection("parcel")
	filter := bson.D{{"parcelid", pid}}
	update := bson.D{
		{"$set",
			bson.D{
				{"status", "delivered"},
			},
		},
	}
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)
	err := parcelColl.FindOneAndUpdate(ctx, filter, update, opts).Decode(parcel)

	return err
}
