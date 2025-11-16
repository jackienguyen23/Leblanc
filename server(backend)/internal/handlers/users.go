package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"leblanc/server/internal/db"
	"leblanc/server/internal/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	NameOrEmail string `json:"name"`
	Password    string `json:"password"`
}

func RegisterUser(c *gin.Context) {
	var req registerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	if req.Name == "" || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name, email and password are required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	coll := db.DB.Collection("users")

	lowerName := strings.ToLower(req.Name)
	lowerEmail := strings.ToLower(req.Email)

	// ensure user doesn't already exist
	filter := bson.M{"$or": []bson.M{
		{"nameLower": lowerName},
		{"emailLower": lowerEmail},
	}}
	if err := coll.FindOne(ctx, filter).Err(); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	} else if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
		return
	}

	user := models.User{
		ID:           primitive.NewObjectID(),
		Name:         req.Name,
		NameLower:    lowerName,
		Email:        req.Email,
		EmailLower:   lowerEmail,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}

	if _, err := coll.InsertOne(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ok": true, "user": user.Public()})
}

func LoginUser(c *gin.Context) {
	var req loginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.NameOrEmail = strings.TrimSpace(req.NameOrEmail)
	req.Password = strings.TrimSpace(req.Password)
	if req.NameOrEmail == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and password are required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	coll := db.DB.Collection("users")

	lookup := strings.ToLower(req.NameOrEmail)
	filter := bson.M{
		"$or": []bson.M{
			{"nameLower": lookup},
			{"emailLower": lookup},
		},
	}

	var user models.User
	if err := coll.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "user": user.Public()})
}
