package handlers

import (
	"context"
	"net/http"
	"time"

	"leblanc/server/internal/db"
	"leblanc/server/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	var b models.Booking
	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := db.DB.Collection("bookings").InsertOne(ctx, b)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, gin.H{"ok": true, "id": res.InsertedID})
}
