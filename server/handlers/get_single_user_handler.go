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

// Get all users
func GetUser(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("In get user handler")
		// Get the user ID parameter
		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Find the user document
		usersCollection := client.Database("programmersdb").Collection("users")
		result := usersCollection.FindOne(context.Background(), bson.M{"_id": id})
		if result.Err() != nil {
			c.AbortWithError(http.StatusNotFound, result.Err())
			return
		}

		// Decode the document and send the response
		var user models.User
		err = result.Decode(&user)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
