package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"CRUD_API/models"
)

// Create a new user
func CreateUser(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("In create user handler")
		// Decode the JSON request body
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Insert the user document
		usersCollection := client.Database("programmersdb").Collection("users")
		result, err := usersCollection.InsertOne(context.Background(), user)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Set the inserted document ID as the user ID
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			user.ID = oid
		}

		// Send the response
		c.JSON(http.StatusCreated, user)
	}
}
