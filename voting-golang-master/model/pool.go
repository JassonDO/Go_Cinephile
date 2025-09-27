package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pool struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PoolId      int64              `bson:"pool_id" json:"pool_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	ClosedAt    int64              `bson:"closed_at" json:"closed_at"`
	Timestamp   int64              `bson:"timestamp" json:"timestamp"`
	Status      int                `bson:"status" json:"status" default:"0"`
}
