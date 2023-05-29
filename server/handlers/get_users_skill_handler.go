package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"CRUD_API/models"
)

// Get users by skill
func GetUsersBySkill(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the skill query parameter
		skill := strings.ToLower(c.Param("skill"))

		// Query the users collection for users with the specified skill
		usersCollection := client.Database("programmersdb").Collection("users")
		filter := bson.M{
			"skills": bson.M{
				"$regex":   fmt.Sprintf("(?i)%s", skill),
				"$options": "i",
			},
		}
		cursor, err := usersCollection.Find(context.Background(), filter)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer cursor.Close(context.Background())

		// Iterate over the cursor and collect the users
		var users []models.User
		for cursor.Next(context.Background()) {
			var user models.User
			err := cursor.Decode(&user)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			// Check if the user's skills list contains the specified skill
			for _, userSkill := range user.Skills {
				if strings.ToLower(userSkill) == skill {
					users = append(users, user)
					break
				}
			}
		}
		if err := cursor.Err(); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Send the response
		c.JSON(http.StatusOK, users)
	}
}
