package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name         string             `bson:"name" json:"name"`
	NameLower    string             `bson:"nameLower" json:"-"`
	Email        string             `bson:"email" json:"email"`
	EmailLower   string             `bson:"emailLower" json:"-"`
	PasswordHash string             `bson:"passwordHash" json:"-"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
}

type PublicUser struct {
	ID        primitive.ObjectID `json:"_id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"createdAt"`
}

func (u User) Public() PublicUser {
	return PublicUser{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
