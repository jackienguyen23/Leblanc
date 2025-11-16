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
    "github.com/graphql-go/graphql"
    ghandler "github.com/graphql-go/handler"
    "go.mongodb.org/mongo-driver/bson"
)

// GraphQL schema for recommendations
func buildRecoSchema() (graphql.Schema, error) {
    emotionFitType := graphql.NewObject(graphql.ObjectConfig{
        Name: "EmotionFit",
        Fields: graphql.Fields{
            "calm":        &graphql.Field{Type: graphql.Float},
            "happy":       &graphql.Field{Type: graphql.Float},
            "stressed":    &graphql.Field{Type: graphql.Float},
            "sad":         &graphql.Field{Type: graphql.Float},
            "adventurous": &graphql.Field{Type: graphql.Float},
        },
    })

    drinkType := graphql.NewObject(graphql.ObjectConfig{
        Name: "Drink",
        Fields: graphql.Fields{
            "_id":      &graphql.Field{Type: graphql.String},
            "name":     &graphql.Field{Type: graphql.String},
            "price":    &graphql.Field{Type: graphql.Int},
            "tags":     &graphql.Field{Type: graphql.NewList(graphql.String)},
            "caffeine": &graphql.Field{Type: graphql.String},
            "temp":     &graphql.Field{Type: graphql.String},
            "sweetness": &graphql.Field{Type: graphql.Int},
            "colorTone": &graphql.Field{Type: graphql.String},
            "emotionFit": &graphql.Field{Type: emotionFitType},
            "image":    &graphql.Field{Type: graphql.String},
            "desc":     &graphql.Field{Type: graphql.String},
            "score":    &graphql.Field{Type: graphql.Float},
        },
    })

    fields := graphql.Fields{
        "reco": &graphql.Field{
            Type: graphql.NewList(drinkType),
            Args: graphql.FieldConfigArgument{
                "emotion":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
                "colorTone": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
                "timeOfDay": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
                "tempPref":  &graphql.ArgumentConfig{Type: graphql.String},
                "topK":      &graphql.ArgumentConfig{Type: graphql.Int},
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                emotion := p.Args["emotion"].(string)
                colorTone := p.Args["colorTone"].(string)
                timeOfDay := p.Args["timeOfDay"].(string)
                var tempPref *string
                if v, ok := p.Args["tempPref"].(string); ok && v != "" {
                    tempPref = &v
                }
                topK := 5
                if v, ok := p.Args["topK"].(int); ok && v > 0 { topK = v }

                payload := services.RecoPayload{
                    Emotion:   emotion,
                    ColorTone: colorTone,
                    Context: services.Context{
                        TimeOfDay: timeOfDay,
                        TempPref:  tempPref,
                    },
                }

                ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
                defer cancel()

                cur, err := db.DB.Collection("drinks").Find(ctx, bson.D{})
                if err != nil { return nil, err }
                var drinks []models.Drink
                if err := cur.All(ctx, &drinks); err != nil { return nil, err }

                type item struct{ D models.Drink; S float64 }
                var ranked []item
                for _, d := range drinks {
                    ranked = append(ranked, item{D: d, S: services.ScoreDrink(d, payload)})
                }
                sort.Slice(ranked, func(i, j int) bool { return ranked[i].S > ranked[j].S })
                if len(ranked) > topK { ranked = ranked[:topK] }

                // build response list of maps that include score
                out := make([]map[string]interface{}, len(ranked))
                for i, r := range ranked {
                    m := map[string]interface{}{}
                    // reuse JSON tags from models by manual mapping
                    m["_id"] = r.D.ID.Hex()
                    m["name"] = r.D.Name
                    m["price"] = r.D.Price
                    m["tags"] = r.D.Tags
                    m["caffeine"] = r.D.Caffeine
                    m["temp"] = r.D.Temp
                    m["sweetness"] = r.D.Sweetness
                    m["colorTone"] = r.D.ColorTone
                    m["emotionFit"] = r.D.EmotionFit
                    m["image"] = r.D.Image
                    m["desc"] = r.D.Desc
                    m["score"] = float64(int(r.S*1000)) / 1000.0
                    out[i] = m
                }

                return out, nil
            },
        },
    }

    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
    return graphql.NewSchema(schemaConfig)
}

// GraphQLHandler returns a gin.HandlerFunc which wraps the graphql-go http handler
func GraphQLHandler() gin.HandlerFunc {
    schema, err := buildRecoSchema()
    if err != nil {
        // if schema build fails, return a handler that reports error
        return func(c *gin.Context) { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) }
    }
    h := ghandler.New(&ghandler.Config{
        Schema:   &schema,
        Pretty:   true,
        GraphiQL: true, // serve GraphiQL playground on GET
    })

    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}
