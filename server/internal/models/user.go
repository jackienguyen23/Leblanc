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
	Role         string             `bson:"role,omitempty" json:"role,omitempty"`
	Verified     bool               `bson:"verified,omitempty" json:"verified,omitempty"`
	PasswordHash string             `bson:"passwordHash" json:"-"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
}

type PublicUser struct {
	ID        primitive.ObjectID `json:"_id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Role      string             `json:"role,omitempty"`
	Verified  bool               `json:"verified"`
	CreatedAt time.Time          `json:"createdAt"`
}

func (u User) Public() PublicUser {
	return PublicUser{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		Verified:  u.Verified,
		CreatedAt: u.CreatedAt,
	}
}
