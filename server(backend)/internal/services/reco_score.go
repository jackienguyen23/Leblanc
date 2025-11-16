package services

import "leblanc/server/internal/models"

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
