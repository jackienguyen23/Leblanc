package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func Init() {
	_ = godotenv.Load()

	uri := os.Getenv("MONGO_URI")
	name := os.Getenv("MONGO_DB")
	// Keep connect attempt short to fail fast if Mongo is unreachable.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Mongo connect error:", err)
	}
	if err := cl.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo ping error:", err)
	}
	Client = cl
	DB = cl.Database(name)
	log.Println("Mongo connected:", uri, "db:", name)
}