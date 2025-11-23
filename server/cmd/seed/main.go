package main

import (
	"context"
	"log"
	"time"

	"leblanc/server/internal/db"
	"leblanc/server/internal/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	_ = godotenv.Load()
	db.Init()

	// Use one shared context so index creation and seeding can't hang forever.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := ensureUserIndexes(ctx); err != nil {
		log.Fatalf("ensure indexes: %v", err)
	}
	if err := seedDrinks(ctx); err != nil {
		log.Fatalf("seed drinks: %v", err)
	}

	log.Println("Database ready.")
}

func ensureUserIndexes(ctx context.Context) error {
	coll := db.DB.Collection("users")
	// Enforce uniqueness for login/lookup fields.
	models := []mongo.IndexModel{
		{Keys: bson.D{{Key: "nameLower", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "emailLower", Value: 1}}, Options: options.Index().SetUnique(true)},
	}
	_, err := coll.Indexes().CreateMany(ctx, models)
	return err
}

func seedDrinks(ctx context.Context) error {
	coll := db.DB.Collection("drinks")
	count, err := coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Printf("drinks collection already has %d documents, skipping seed", count)
		return nil
	}

	// Curated list that powers the menu and recommendations.
	drinks := []models.Drink{
		{
			ID:        primitive.NewObjectID(),
			Name:      "Moonlit Lavender Latte",
			Price:     65000,
			Tags:      []string{"signature", "floral"},
			Caffeine:  "med",
			Temp:      "either",
			Sweetness: 3,
			ColorTone: "cool",
			EmotionFit: models.EmotionFit{
				Calm: 0.9, Happy: 0.6, Stressed: 0.85, Sad: 0.7, Adventurous: 0.4,
			},
			Image: "https://images.unsplash.com/photo-1512568400610-62da28bc8a13?auto=format&fit=crop&w=720&q=80",
			Desc:  "Lavender infused espresso with steamed oat milk and vanilla bean syrup.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Amber Salted Caramel Cold Brew",
			Price:     59000,
			Tags:      []string{"cold brew", "caramel"},
			Caffeine:  "high",
			Temp:      "iced",
			Sweetness: 2,
			ColorTone: "warm",
			EmotionFit: models.EmotionFit{
				Calm: 0.5, Happy: 0.8, Stressed: 0.6, Sad: 0.55, Adventurous: 0.6,
			},
			Image: "https://images.unsplash.com/photo-1464306076886-da185f6a9d12?auto=format&fit=crop&w=720&q=80",
			Desc:  "12-hour cold brew finished with salted caramel foam and cacao nib dust.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Golden Turmeric Oat Cortado",
			Price:     62000,
			Tags:      []string{"oat", "spiced"},
			Caffeine:  "low",
			Temp:      "hot",
			Sweetness: 1,
			ColorTone: "warm",
			EmotionFit: models.EmotionFit{
				Calm: 0.8, Happy: 0.6, Stressed: 0.7, Sad: 0.9, Adventurous: 0.5,
			},
			Image: "https://images.unsplash.com/photo-1432107294467-6af17c6a1c67?auto=format&fit=crop&w=720&q=80",
			Desc:  "Turmeric, ginger, and local honey layered with ristretto shots and micro-foamed oat milk.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Citrus Tonic Espresso",
			Price:     55000,
			Tags:      []string{"spritz", "espresso"},
			Caffeine:  "med",
			Temp:      "iced",
			Sweetness: 2,
			ColorTone: "neutral",
			EmotionFit: models.EmotionFit{
				Calm: 0.4, Happy: 0.85, Stressed: 0.4, Sad: 0.5, Adventurous: 0.9,
			},
			Image: "https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=720&q=80",
			Desc:  "Double espresso poured over yuzu tonic with grapefruit zest and rosemary mist.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Velvet Midnight Mocha",
			Price:     67000,
			Tags:      []string{"mocha", "dark chocolate"},
			Caffeine:  "med",
			Temp:      "either",
			Sweetness: 4,
			ColorTone: "cool",
			EmotionFit: models.EmotionFit{
				Calm: 0.7, Happy: 0.65, Stressed: 0.75, Sad: 0.8, Adventurous: 0.5,
			},
			Image: "https://images.unsplash.com/photo-1498804103079-a6351b050096?auto=format&fit=crop&w=720&q=80",
			Desc:  "Single-origin dark chocolate melted with espresso, topped with smoked sea salt cream.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Matcha Cascara Fizz",
			Price:     60000,
			Tags:      []string{"matcha", "tea"},
			Caffeine:  "low",
			Temp:      "iced",
			Sweetness: 3,
			ColorTone: "neutral",
			EmotionFit: models.EmotionFit{
				Calm: 0.75, Happy: 0.7, Stressed: 0.65, Sad: 0.6, Adventurous: 0.7,
			},
			Image: "https://images.unsplash.com/photo-1422207239328-29f83822d1f2?auto=format&fit=crop&w=720&q=80",
			Desc:  "Ceremonial matcha shaken with cascara reduction and sparkling mineral water.",
		},
	}

	docs := make([]any, len(drinks))
	for i, d := range drinks {
		docs[i] = d
	}

	_, err = coll.InsertMany(ctx, docs)
	return err
}
