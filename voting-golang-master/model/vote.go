package model

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vote struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PoolId     int64              `bson:"pool_id" json:"pool_id"`
	CadidateId int64              `bson:"candidate_id" json:"candidate_id"`
	Name       string             `bson:"name" json:"name"`
	Voter      string             `bson:"voter" json:"voter"`
	Timestamp  int64              `bson:"timestamp" json:"timestamp"`

	Pool Pool `json:"pool"`
	User User `json:"user"`
}

func (v Vote) ToMap() (map[string]interface{}, error) {

	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil

}
