package services

import (
	"math"
	"sort"

	"leblanc/server/internal/models"
)

// payload structures for reco
type Context struct {
	TimeOfDay string  `json:"timeOfDay"` // "day"|"night"
	TempPref  *string `json:"tempPref"`  // optional: "hot"|"iced"
}
type RecoPayload struct {
	Emotion   string  `json:"emotion"`
	ColorTone string  `json:"colorTone"`
	Context   Context `json:"context"`
}

func ScoreDrink(d models.Drink, p RecoPayload) float64 {
	// emotion
	em := map[string]float64{
		"calm": d.EmotionFit.Calm, "happy": d.EmotionFit.Happy,
		"stressed": d.EmotionFit.Stressed, "sad": d.EmotionFit.Sad,
		"adventurous": d.EmotionFit.Adventurous,
	}
	sEmotion := em[p.Emotion]
	if sEmotion == 0 { sEmotion = 0.5 }

	// color tone
	sColor := 0.2
	if d.ColorTone == p.ColorTone { sColor = 1.0
	} else if d.ColorTone == "neutral" || p.ColorTone == "neutral" { sColor = 0.6 }

	// context
	sContext := 0.5
	if p.Context.TimeOfDay == "night" && d.Temp == "hot" { sContext += 0.2 }
	if p.Context.TimeOfDay == "day" && d.Temp == "iced" { sContext += 0.2 }
	if p.Context.TempPref != nil && (d.Temp == *p.Context.TempPref || d.Temp == "either") {
		sContext += 0.3
	}
	return 0.5*sEmotion + 0.3*sColor + 0.2*sContext
}

// DrinkScore represents a drink with its recommendation score
type DrinkScore struct {
	DrinkID string
	Score   float64
}

// ScoreDrinks scores drinks based on emotion fit and optional preferences
func ScoreDrinks(drinks []models.Drink, emotionFit models.EmotionFit, caffeine, temp string, sweetness int) []DrinkScore {
	var scores []DrinkScore

	for _, drink := range drinks {
		// Calculate emotion fit score using cosine similarity
		emotionScore := calculateEmotionScore(drink.EmotionFit, emotionFit)
		
		// Initialize preference score
		prefScore := 1.0
		
		// Apply caffeine preference if specified
		if caffeine != "" && drink.Caffeine != caffeine {
			prefScore *= 0.5
		}
		
		// Apply temperature preference if specified
		if temp != "" && drink.Temp != "either" && drink.Temp != temp {
			prefScore *= 0.5
		}
		
		// Apply sweetness preference if specified
		if sweetness > 0 {
			sweetnessDiff := math.Abs(float64(drink.Sweetness - sweetness))
			sweetnessScore := 1.0 - (sweetnessDiff / 10.0)
			if sweetnessScore < 0 {
				sweetnessScore = 0
			}
			prefScore *= (0.5 + 0.5*sweetnessScore)
		}
		
		// Combine emotion score and preference score
		totalScore := emotionScore * 0.7 + prefScore * 0.3
		
		scores = append(scores, DrinkScore{
			DrinkID: drink.ID.Hex(),
			Score:   math.Round(totalScore*1000) / 1000,
		})
	}
	
	// Sort by score descending
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})
	
	// Return top 5
	if len(scores) > 5 {
		scores = scores[:5]
	}
	
	return scores
}

// calculateEmotionScore computes cosine similarity between two emotion vectors
func calculateEmotionScore(e1, e2 models.EmotionFit) float64 {
	// Create vectors
	v1 := []float64{e1.Calm, e1.Happy, e1.Stressed, e1.Sad, e1.Adventurous}
	v2 := []float64{e2.Calm, e2.Happy, e2.Stressed, e2.Sad, e2.Adventurous}
	
	// Calculate dot product
	dotProduct := 0.0
	for i := range v1 {
		dotProduct += v1[i] * v2[i]
	}
	
	// Calculate magnitudes
	mag1 := 0.0
	mag2 := 0.0
	for i := range v1 {
		mag1 += v1[i] * v1[i]
		mag2 += v2[i] * v2[i]
	}
	mag1 = math.Sqrt(mag1)
	mag2 = math.Sqrt(mag2)
	
	// Avoid division by zero
	if mag1 == 0 || mag2 == 0 {
		return 0
	}
	
	// Return cosine similarity
	return dotProduct / (mag1 * mag2)
}
