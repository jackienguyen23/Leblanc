package main

import (
	"log"
	"os"

	"leblanc/server/internal/db"
	"leblanc/server/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db.Init()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"msg": "LeBlanc Go API ✅"}) })
	r.GET("/drinks", handlers.GetDrinks)
	r.POST("/reco/from-features", handlers.RecoFromFeatures)
	r.POST("/bookings", handlers.CreateBooking)

	port := os.Getenv("PORT")
	if port == "" { port = "4000" }
	log.Println("Server listening on http://localhost:" + port)
	_ = r.Run(":" + port)
}
