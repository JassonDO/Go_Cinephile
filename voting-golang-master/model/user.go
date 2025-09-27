package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Address string             `bson:"address" json:"address"`
	Email   string             `bson:"email" json:"email"`
	Name    string             `bson:"name" json:"name"`
}
