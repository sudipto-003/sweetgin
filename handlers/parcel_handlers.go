package handlers

import (
	"context"

	"github.com/sudipto-003/sweet-gin/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ParcelCollection struct {
	Collection *mongo.Collection
}

func (mcoll *ParcelCollection) CreateNewParcelOreder(ctx context.Context, parcel *models.Parcel) error {
	res, err := mcoll.Collection.InsertOne(ctx, parcel)
	if err != nil {
		return err
	}

	parcel.ParcelId = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}
