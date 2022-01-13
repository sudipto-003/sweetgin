package models

type Parcel struct {
	ParcelId   string        `json:"_id" binding:"-"`
	Weight     float32       `json:"weight"`
	Status     string        `json:"status"`
	PickedDate string        `json:"date_picked"`
	Sender     AddressDetail `json:"sender"`
	Receiver   AddressDetail `json:"receiver"`
}

type AddressDetail struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	ZIPCode string `json:"zip_code"`
}
