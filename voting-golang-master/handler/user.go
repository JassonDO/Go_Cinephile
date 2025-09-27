package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"voting-golang/config"
	"voting-golang/model"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserHandler(r *chi.Mux) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/{address}", GetUserByAddress)
		r.Put("/{address}", UpdateUser)
	})
}

func UpdateUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	address := chi.URLParam(r, "address")

	db := config.MongoClient.Database("voting")

	collection := db.Collection("user")

	var body UpdateUserRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	opts := options.Update().SetUpsert(true)

	_, err = collection.UpdateOne(r.Context(), bson.M{"address": address}, bson.M{"$set": bson.M{
		"name":  body.Name,
		"email": body.Email,
	}}, opts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user model.User
	err = collection.FindOne(r.Context(), bson.M{"address": address}).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)

}

func GetUserByAddress(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	db := config.MongoClient.Database("voting")

	collection := db.Collection("user")

	var user model.User
	err := collection.FindOne(r.Context(), bson.M{"address": address}).Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			json.NewEncoder(w).Encode(user)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
