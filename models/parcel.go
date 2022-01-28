package models

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type pTime primitive.DateTime

const timeFormat = "02-01-2006"

func (pt *pTime) UnmarshalJSON(bs []byte) error {
	var timestr string
	err := json.Unmarshal(bs, &timestr)
	if err != nil {
		return err
	}

	ptime, err := time.Parse(timeFormat, timestr)
	*pt = pTime(primitive.NewDateTimeFromTime(ptime))
	return err
}

func (pt pTime) MarshalJSON() ([]byte, error) {
	timestr := primitive.DateTime(pt).Time().Format(timeFormat)
	return json.Marshal(timestr)
}

type Parcel struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id" binding:"-"`
	ParcelId   string             `json:"parcel_id"`
	Weight     float32            `json:"weight"`
	Status     string             `json:"status"`
	PickedDate pTime              `json:"date_picked" bson:"date_picked"`
	Sender     AddressDetail      `json:"sender"`
	Receiver   AddressDetail      `json:"receiver"`
}

type AddressDetail struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	ZIPCode string `json:"zip_code"`
}
