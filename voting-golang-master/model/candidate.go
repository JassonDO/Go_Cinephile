package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Candidate struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PoolId      int64              `bson:"pool_id" json:"pool_id"`
	CandidateId int64              `bson:"candidate_id" json:"candidate_id"`
	Name        string             `bson:"name" json:"name"`
	VoteCount   int64              `bson:"vote_count" json:"vote_count"`
}
