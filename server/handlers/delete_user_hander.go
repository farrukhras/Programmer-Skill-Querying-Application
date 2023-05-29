package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// delete a user
func DeleteUser(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("In delete user handler")
		// Get the user ID parameter
		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Delete the user document
		usersCollection := client.Database("programmersdb").Collection("users")
		result, err := usersCollection.DeleteOne(context.Background(), bson.M{"_id": id})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Check if the document was deleted
		if result.DeletedCount == 0 {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		// Send the response
		c.Status(http.StatusNoContent)
	}
}
