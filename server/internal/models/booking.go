package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingItem struct {
	DrinkID primitive.ObjectID `bson:"drinkId" json:"drinkId"`
	Qty     int                `bson:"qty" json:"qty"`
	Options map[string]any     `bson:"options" json:"options"`
}

type Booking struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name    string             `bson:"name" json:"name"`
	Phone   string             `bson:"phone" json:"phone"`
	Time    time.Time          `bson:"time" json:"time"`
	Items   []BookingItem      `bson:"items" json:"items"`
	Channel string             `bson:"channel" json:"channel"`
}
