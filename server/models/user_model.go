package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Skills   []string           `json:"skills,omitempty" bson:"skills,omitempty"`
	Image    string             `json:"image,omitempty" bson:"image,omitempty"`
	JobTitle string             `json:"jobTitle,omitempty" bson:"jobTitle,omitempty"`
}
