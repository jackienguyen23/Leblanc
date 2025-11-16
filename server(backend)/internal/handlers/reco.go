package handlers

import (
	"context"
	"net/http"
	"sort"
	"time"

	"leblanc/server/internal/db"
	"leblanc/server/internal/models"
	"leblanc/server/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func RecoFromFeatures(c *gin.Context) {
	var payload services.RecoPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := db.DB.Collection("drinks").Find(ctx, bson.D{})
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }

	var drinks []models.Drink
	if err := cur.All(ctx, &drinks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return
	}

	type item struct{ D models.Drink; S float64 }
	var ranked []item
	for _, d := range drinks {
		ranked = append(ranked, item{D: d, S: services.ScoreDrink(d, payload)})
	}
	sort.Slice(ranked, func(i, j int) bool { return ranked[i].S > ranked[j].S })
	if len(ranked) > 5 { ranked = ranked[:5] }

	// attach score to response
	type out struct {
		models.Drink
		Score float64 `json:"score"`
	}
	resp := make([]out, len(ranked))
	for i, r := range ranked {
		resp[i] = out{Drink: r.D, Score: float64(int(r.S*1000)) / 1000.0}
	}
	c.JSON(http.StatusOK, resp)
}
