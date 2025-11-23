package main

import (
	"log"
	"os"

	"leblanc/server/internal/db"
	"leblanc/server/internal/graph"
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

	// REST API endpoints
	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"msg": "LeBlanc Go API with REST & GraphQL."}) })
	r.GET("/drinks", handlers.GetDrinks)
	r.POST("/reco/from-features", handlers.RecoFromFeatures)
	r.POST("/bookings", handlers.CreateBooking)
	r.POST("/auth/register", handlers.RegisterUser)
	r.POST("/auth/login", handlers.LoginUser)

	// GraphQL endpoint
	r.POST("/graphql", graph.Handler())
	r.GET("/graphql", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "GraphQL endpoint - send POST requests with GraphQL queries"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	log.Println("Server listening on http://localhost:" + port)
	log.Println("REST API: http://localhost:" + port)
	log.Println("GraphQL: http://localhost:" + port + "/graphql")
	_ = r.Run(":" + port)
}
