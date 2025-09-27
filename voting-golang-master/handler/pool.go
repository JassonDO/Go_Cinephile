package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
	"voting-golang/config"
	"voting-golang/model"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPoolHandler(r *chi.Mux) {
	r.Route("/pool", func(r chi.Router) {
		r.Get("/", GetPool)
		r.Get("/{id}", GetPoolById)
		r.Get("/{id}/winner", GetWinner)
	})
}

func GetPool(w http.ResponseWriter, r *http.Request) {
	db := config.MongoClient.Database("voting")
	collection := db.Collection("pool")
	cursor, err := collection.Find(r.Context(), bson.M{})
	if err != nil {
		log.Println(err)
	}

	var pools []model.Pool
	if err = cursor.All(r.Context(), &pools); err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pools)

}

func GetPoolById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := config.MongoClient.Database("voting")
	collection := db.Collection("pool")
	cursor := collection.FindOne(r.Context(), bson.M{"pool_id": id})

	if cursor.Err() != nil {
		if cursor.Err() == mongo.ErrNoDocuments {
			http.Error(w, "Pool not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, cursor.Err().Error(), http.StatusInternalServerError)
			return
		}

	}

	var pool model.Pool
	if err := cursor.Decode(&pool); err != nil {
		http.Error(w, cursor.Err().Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pool)

}

func GetWinner(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	winner, err := GetWinnerByPoolId(r.Context(), int64(id))
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(winner)

}

func GetWinnerByPoolId(ctx context.Context, id int64) ([]model.Candidate, error) {
	db := config.MongoClient.Database("voting")
	poolCollection := db.Collection("pool")
	candidateCollection := db.Collection("candidate")

	poolCursor := poolCollection.FindOne(ctx, bson.M{"pool_id": id})
	if poolCursor.Err() != nil {
		return nil, poolCursor.Err()

	}

	var pool model.Pool
	if err := poolCursor.Decode(&pool); err != nil {
		return nil, poolCursor.Err()
	}

	if pool.ClosedAt > time.Now().Unix() {
		return nil, errors.New("Pool is not closed yet")
	}

	candidateCursor, err := candidateCollection.Find(ctx, bson.M{"pool_id": id})

	if err != nil {
		return nil, err
	}

	var candidates []model.Candidate
	if err = candidateCursor.All(ctx, &candidates); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}

	if len(candidates) == 0 {
		return []model.Candidate{}, nil
	}

	var candidatesCount = make(map[int64][]model.Candidate)
	for _, candidate := range candidates {
		candidatesCount[candidate.VoteCount] = append(candidatesCount[candidate.VoteCount], candidate)
	}
	max := int64(0)
	for count := range candidatesCount {
		if count > max {
			max = count
		}
	}

	if max <= 0 {
		return []model.Candidate{}, nil
	}

	winner := candidatesCount[max]
	return winner, nil
}
