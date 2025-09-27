package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"voting-golang/config"
	"voting-golang/model"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewVoteHandler(r *chi.Mux) {
	r.Route("/vote", func(r chi.Router) {
		r.Get("/", GetAllVote)
		r.Get("/{voter}/histories", GetVoteHistories)
		r.Get("/{voter}/{pool_id}", getVoteByVoterAndPoolId)
	})
}

func GetAllVote(w http.ResponseWriter, r *http.Request) {
	db := config.MongoClient.Database("voting")

	collection := db.Collection("vote")
	cursor, err := collection.Find(r.Context(), bson.M{})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var votes []model.Vote
	if err := cursor.All(context.Background(), &votes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(votes)

}

func GetVoteHistories(w http.ResponseWriter, r *http.Request) {
	voter := chi.URLParam(r, "voter")

	db := config.MongoClient.Database("voting")

	collection := db.Collection("vote")
	cursor, err := collection.Aggregate(r.Context(), mongo.Pipeline{
		bson.D{{"$match", bson.D{{"voter", voter}}}},
		bson.D{{"$lookup", bson.D{
			{"from", "pool"},
			{"localField", "pool_id"},
			{"foreignField", "pool_id"},
			{"as", "pool"},
		}}},
		bson.D{{"$unwind", bson.D{
			{"path", "$pool"},
		}}},
	})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			json.NewEncoder(w).Encode([]model.Vote{})
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var voteHistories []model.Vote
	for cursor.Next(context.Background()) {
		var vote model.Vote
		cursor.Decode(&vote)
		voteHistories = append(voteHistories, vote)
	}

	json.NewEncoder(w).Encode(voteHistories)

}

func getVoteByVoterAndPoolId(w http.ResponseWriter, r *http.Request) {
	voter := chi.URLParam(r, "voter")
	poolId, err := strconv.Atoi(chi.URLParam(r, "pool_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := config.MongoClient.Database("voting")

	collection := db.Collection("vote")
	var vote model.Vote
	err = collection.FindOne(context.Background(), bson.M{"voter": voter, "pool_id": poolId}).Decode(&vote)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			json.NewEncoder(w).Encode(model.Vote{})
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(vote)
}
