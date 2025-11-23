package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

type GraphQLResponse struct {
	Data   interface{} `json:"data,omitempty"`
	Errors []string    `json:"errors,omitempty"`
}

// Handler creates a Gin handler for GraphQL
func Handler() gin.HandlerFunc {
	resolver := &Resolver{}
	
	return func(c *gin.Context) {
		var req GraphQLRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, GraphQLResponse{
				Errors: []string{err.Error()},
			})
			return
		}

		ctx := context.Background()
		result, err := executeQuery(ctx, resolver, req.Query, req.Variables)
		if err != nil {
			c.JSON(http.StatusOK, GraphQLResponse{
				Errors: []string{err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, GraphQLResponse{
			Data: result,
		})
	}
}

// Simple query executor (this is a simplified version - in production, use a proper GraphQL library)
func executeQuery(ctx context.Context, resolver *Resolver, query string, variables map[string]interface{}) (interface{}, error) {
	// This is a simplified implementation
	// In a real application, you would use gqlgen or graphql-go to parse and execute queries
	
	// For now, we'll create a basic routing based on query content
	if contains(query, "query") {
		if contains(query, "drinks") {
			return resolver.Drinks(ctx)
		}
		if contains(query, "users") {
			return resolver.Users(ctx)
		}
		if contains(query, "bookings") {
			return resolver.Bookings(ctx)
		}
	}
	
	if contains(query, "mutation") {
		if contains(query, "createBooking") {
			var input CreateBookingInput
			if inputData, ok := variables["input"].(map[string]interface{}); ok {
				jsonData, _ := json.Marshal(inputData)
				json.Unmarshal(jsonData, &input)
				return resolver.CreateBooking(ctx, input)
			}
		}
		if contains(query, "register") {
			var input RegisterInput
			if inputData, ok := variables["input"].(map[string]interface{}); ok {
				jsonData, _ := json.Marshal(inputData)
				json.Unmarshal(jsonData, &input)
				return resolver.Register(ctx, input)
			}
		}
		if contains(query, "login") {
			var input LoginInput
			if inputData, ok := variables["input"].(map[string]interface{}); ok {
				jsonData, _ := json.Marshal(inputData)
				json.Unmarshal(jsonData, &input)
				return resolver.Login(ctx, input)
			}
		}
		if contains(query, "recommendFromFeatures") {
			var emotionFit EmotionFitInput
			if emotionData, ok := variables["emotionFit"].(map[string]interface{}); ok {
				jsonData, _ := json.Marshal(emotionData)
				json.Unmarshal(jsonData, &emotionFit)
			}
			
			var caffeine, temp *string
			var sweetness *int
			
			if c, ok := variables["caffeine"].(string); ok {
				caffeine = &c
			}
			if t, ok := variables["temp"].(string); ok {
				temp = &t
			}
			if s, ok := variables["sweetness"].(float64); ok {
				sInt := int(s)
				sweetness = &sInt
			}
			
			return resolver.RecommendFromFeatures(ctx, emotionFit, caffeine, temp, sweetness)
		}
	}
	
	return nil, fmt.Errorf("unsupported query")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && 
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || 
		findInString(s, substr)))
}

func findInString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
