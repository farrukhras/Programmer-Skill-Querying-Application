package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"CRUD_API/models"
)

// Get all users
func GetUsers(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("In get all users handler ")
		// Get a handle to the users collection
		usersCollection := client.Database("programmersdb").Collection("users")

		// Find all documents in the collection
		cur, err := usersCollection.Find(context.Background(), bson.D{})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer cur.Close(context.Background())

		// Iterate over the documents and add them to a slice
		var users []models.User
		for cur.Next(context.Background()) {
			var user models.User
			err := cur.Decode(&user)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			users = append(users, user)
		}

		// Send the response
		c.JSON(http.StatusOK, users)
	}
}
