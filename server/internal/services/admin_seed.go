package services

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"leblanc/server/internal/db"
	"leblanc/server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// EnsureAdminUser inserts a predefined admin account if it does not already exist.
// Controlled by env: ADMIN_NAME, ADMIN_EMAIL, ADMIN_PASSWORD.
func EnsureAdminUser() {
	name := strings.TrimSpace(os.Getenv("ADMIN_NAME"))
	email := strings.TrimSpace(os.Getenv("ADMIN_EMAIL"))
	password := os.Getenv("ADMIN_PASSWORD")

	if email == "" || password == "" {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	coll := db.DB.Collection("users")
	emailLower := strings.ToLower(email)

	var existing models.User
	err := coll.FindOne(ctx, bson.M{"emailLower": emailLower}).Decode(&existing)
	if err == nil {
		// already exists, ensure role admin
		if existing.Role != "admin" || !existing.Verified {
			_, _ = coll.UpdateOne(ctx, bson.M{"_id": existing.ID}, bson.M{"$set": bson.M{"role": "admin", "verified": true}})
		}
		return
	} else if err != mongo.ErrNoDocuments {
		log.Printf("ensure admin user: lookup error: %v", err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("ensure admin user: hash error: %v", err)
		return
	}

	if name == "" {
		name = "Admin"
	}

	admin := models.User{
		ID:           primitive.NewObjectID(),
		Name:         name,
		NameLower:    strings.ToLower(name),
		Email:        email,
		EmailLower:   emailLower,
		Role:         "admin",
		Verified:     true,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}

	if _, err := coll.InsertOne(ctx, admin); err != nil {
		log.Printf("ensure admin user: insert error: %v", err)
		return
	}

	log.Printf("admin account ensured for %s", email)
}
