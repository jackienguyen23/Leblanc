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

	// New curated list grouped by day (coffee) vs night (pub/cocktail).
	drinks := []models.Drink{
		{
			ID:        primitive.NewObjectID(),
			Name:      "Espresso",
			Price:     35000,
			Tags:      []string{"coffee", "day", "espresso"},
			Caffeine:  "high",
			Temp:      "hot",
			Sweetness: 1,
			ColorTone: "warm",
			EmotionFit: models.EmotionFit{
				Calm: 0.3, Happy: 0.5, Stressed: 0.7, Sad: 0.4, Adventurous: 0.6,
			},
			Image: "https://images.unsplash.com/photo-1510626176961-4b37d0b4e904?auto=format&fit=crop&w=900&q=80",
			Desc:  "Shot of espresso with your choice of bean profile: arabica, robusta, moka, or culi.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Nâu đá",
			Price:     40000,
			Tags:      []string{"coffee", "day", "iced"},
			Caffeine:  "high",
			Temp:      "iced",
			Sweetness: 3,
			ColorTone: "warm",
			EmotionFit: models.EmotionFit{
				Calm: 0.4, Happy: 0.7, Stressed: 0.6, Sad: 0.5, Adventurous: 0.5,
			},
			Image: "https://images.unsplash.com/photo-1509042239860-f550ce710b93?auto=format&fit=crop&w=900&q=80",
			Desc:  "Vietnamese iced coffee with condensed milk. Beans: arabica, robusta, moka, or culi.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Bạc sỉu",
			Price:     38000,
			Tags:      []string{"coffee", "day", "milk-forward"},
			Caffeine:  "med",
			Temp:      "hot",
			Sweetness: 4,
			ColorTone: "neutral",
			EmotionFit: models.EmotionFit{
				Calm: 0.5, Happy: 0.75, Stressed: 0.55, Sad: 0.6, Adventurous: 0.4,
			},
			Image: "https://images.unsplash.com/photo-1512568400610-62da28bc8a13?auto=format&fit=crop&w=900&q=80",
			Desc:  "Creamy Saigon-style coffee heavy on milk. Bean picks: arabica, robusta, moka, culi.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Latte",
			Price:     45000,
			Tags:      []string{"coffee", "day"},
			Caffeine:  "med",
			Temp:      "either",
			Sweetness: 2,
			ColorTone: "warm",
			EmotionFit: models.EmotionFit{
				Calm: 0.6, Happy: 0.65, Stressed: 0.5, Sad: 0.5, Adventurous: 0.4,
			},
			Image: "https://images.unsplash.com/photo-1470337458703-46ad1756a187?auto=format&fit=crop&w=900&q=80",
			Desc:  "Silky espresso with steamed milk. Customize with cốm, vani, sữa đặc, sữa kem, hoặc chuối.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Mocha",
			Price:     48000,
			Tags:      []string{"coffee", "day", "chocolate"},
			Caffeine:  "med",
			Temp:      "either",
			Sweetness: 3,
			ColorTone: "cool",
			EmotionFit: models.EmotionFit{
				Calm: 0.55, Happy: 0.7, Stressed: 0.55, Sad: 0.6, Adventurous: 0.45,
			},
			Image: "https://images.unsplash.com/photo-1459257868276-5e65389e2722?auto=format&fit=crop&w=900&q=80",
			Desc:  "Espresso, cocoa, and velvety milk. Bean options: arabica, robusta, moka, culi.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Macchiato",
			Price:     47000,
			Tags:      []string{"coffee", "day"},
			Caffeine:  "med",
			Temp:      "hot",
			Sweetness: 1,
			ColorTone: "neutral",
			EmotionFit: models.EmotionFit{
				Calm: 0.45, Happy: 0.55, Stressed: 0.65, Sad: 0.4, Adventurous: 0.5,
			},
			Image: "https://images.unsplash.com/photo-1495474472287-4d71bcdd2085?auto=format&fit=crop&w=900&q=80",
			Desc:  "Espresso marked with microfoam. Choose beans and add flavors: vani, sữa kem, chuối, hoặc cốm.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Cappuccino",
			Price:     45000,
			Tags:      []string{"coffee", "day"},
			Caffeine:  "med",
			Temp:      "hot",
			Sweetness: 2,
			ColorTone: "warm",
			EmotionFit: models.EmotionFit{
				Calm: 0.6, Happy: 0.65, Stressed: 0.45, Sad: 0.45, Adventurous: 0.4,
			},
			Image: "https://images.unsplash.com/photo-1485808191679-5f86510681a2?auto=format&fit=crop&w=900&q=80",
			Desc:  "Balanced espresso, steamed milk, and foam. Beans: arabica, robusta, moka, culi.",
		},
		// Night / pub menu
		{
			ID:        primitive.NewObjectID(),
			Name:      "Sangria - Tình đơn phương",
			Price:     120000,
			Tags:      []string{"night", "cocktail", "wine"},
			Caffeine:  "none",
			Temp:      "cold",
			Sweetness: 3,
			ColorTone: "warm",
			EmotionFit: models.EmotionFit{
				Calm: 0.5, Happy: 0.7, Stressed: 0.4, Sad: 0.8, Adventurous: 0.6,
			},
			Image: "https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=900&q=80",
			Desc:  "Red wine, citrus, and stone fruits—gently sweet like một mối tình đơn phương.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Margarita - Rối lòng",
			Price:     130000,
			Tags:      []string{"night", "cocktail", "citrus"},
			Caffeine:  "none",
			Temp:      "cold",
			Sweetness: 2,
			ColorTone: "cool",
			EmotionFit: models.EmotionFit{
				Calm: 0.4, Happy: 0.6, Stressed: 0.7, Sad: 0.5, Adventurous: 0.7,
			},
			Image: "https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=900&q=80&sat=-10",
			Desc:  "Tequila, lime, and salted rim—cú xoáy rối lòng nhưng đầy tỉnh táo.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Mojito - Thanh lọc",
			Price:     110000,
			Tags:      []string{"night", "cocktail", "herbal"},
			Caffeine:  "none",
			Temp:      "cold",
			Sweetness: 3,
			ColorTone: "neutral",
			EmotionFit: models.EmotionFit{
				Calm: 0.7, Happy: 0.75, Stressed: 0.45, Sad: 0.5, Adventurous: 0.6,
			},
			Image: "https://images.unsplash.com/photo-1544145945-19cc90f9c6bf?auto=format&fit=crop&w=900&q=80",
			Desc:  "Rum, bạc hà, lime, và soda—thanh lọc và sảng khoái.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Pimm’s - Ngậm ngùi",
			Price:     120000,
			Tags:      []string{"night", "cocktail", "fruity"},
			Caffeine:  "none",
			Temp:      "cold",
			Sweetness: 2,
			ColorTone: "warm",
			EmotionFit: models.EmotionFit{
				Calm: 0.5, Happy: 0.55, Stressed: 0.5, Sad: 0.7, Adventurous: 0.5,
			},
			Image: "https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=900&q=80&sat=5",
			Desc:  "Pimm’s, dưa leo, táo, gừng ale—ngậm ngùi nhưng dễ uống.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Piña Colada - Tức giận",
			Price:     115000,
			Tags:      []string{"night", "cocktail", "tropical"},
			Caffeine:  "none",
			Temp:      "cold",
			Sweetness: 3,
			ColorTone: "cool",
			EmotionFit: models.EmotionFit{
				Calm: 0.55, Happy: 0.7, Stressed: 0.4, Sad: 0.6, Adventurous: 0.65,
			},
			Image: "https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=900&q=80&sat=-5",
			Desc:  "Rum, dứa, và cream of coconut—tropical nhưng dư vị hơi 'tức giận'.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Caipirinha - Ích kỉ tự trách",
			Price:     120000,
			Tags:      []string{"night", "cocktail", "citrus"},
			Caffeine:  "none",
			Temp:      "cold",
			Sweetness: 2,
			ColorTone: "neutral",
			EmotionFit: models.EmotionFit{
				Calm: 0.45, Happy: 0.6, Stressed: 0.55, Sad: 0.65, Adventurous: 0.7,
			},
			Image: "https://images.unsplash.com/photo-1497534446932-c925b458314e?auto=format&fit=crop&w=900&q=80",
			Desc:  "Cachaça, lime, đường mía—sắc chua gắt như lúc tự trách.",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Negroni - Đắng cay",
			Price:     140000,
			Tags:      []string{"night", "cocktail", "bitter"},
			Caffeine:  "none",
			Temp:      "cold",
			Sweetness: 1,
			ColorTone: "cool",
			EmotionFit: models.EmotionFit{
				Calm: 0.4, Happy: 0.5, Stressed: 0.65, Sad: 0.75, Adventurous: 0.8,
			},
			Image: "https://images.unsplash.com/photo-1470337458703-46ad1756a187?auto=format&fit=crop&w=900&q=80&sat=-10",
			Desc:  "Gin, Campari, vermouth đỏ—đắng cay nhưng sâu sắc, dành cho đêm pub.",
		},
	}

	docs := make([]any, len(drinks))
	for i, d := range drinks {
		docs[i] = d
	}

	_, err = coll.InsertMany(ctx, docs)
	return err
}
