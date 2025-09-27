package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Metadata struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	FirstBlock   int64              `bson:"first_block"`
	CurrentBlock int64              `bson:"current_block"`
}
