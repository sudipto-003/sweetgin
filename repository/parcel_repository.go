package repository

import (
	"context"

	"github.com/sudipto-003/sweet-gin/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (client *Repos) CreateNewParcelOreder(ctx context.Context, parcel *models.Parcel) error {
	parcelColl := client.MongoClient.Database("sweetgin").Collection("parcel")
	res, err := parcelColl.InsertOne(ctx, parcel)
	if err != nil {
		return err
	}

	parcel.ParcelId = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}
