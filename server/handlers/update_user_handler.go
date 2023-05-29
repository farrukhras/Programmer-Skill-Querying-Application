package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"CRUD_API/models"
)

// update a user
func UpdateUser(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("In update user handler")
		// Get the user ID parameter
		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Decode the JSON request body
		var user models.User
		err = c.BindJSON(&user)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Update the user document
		usersCollection := client.Database("programmersdb").Collection("users")
		result, err := usersCollection.ReplaceOne(context.Background(), bson.M{"_id": id}, user)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Check if the document was updated
		if result.MatchedCount == 0 {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		// Send the response
		c.Status(http.StatusOK)
	}
}
