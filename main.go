package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"rest-go/helper"
	"rest-go/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = helper.ConnectDb()

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "apllication/json")

	var users []models.Users
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.GetError(err, w)
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var user models.Users

		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(users)

}

func getUsersbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "apllication/json")

	var user models.Users
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{
		"_id": id,
	}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		helper.GetError(err, w)
	}
	json.NewEncoder(w).Encode(user)
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user models.Users

	_ = json.NewDecoder(r.Body).Decode(&user)

	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user models.Users

	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&user)

	update := bson.D{{Key:"$set", Value :bson.D{
		{Key :"name", Value: user.Name},
		{Key :"Dob",Value: user.Dob},
		{Key:"Address",Value: user.Address},
		{Key :"Description",Value:  user.Desription},
		{Key:"CreatedAt",Value: user.CreatedAt},
	}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	user.ID = id
	json.NewEncoder(w).Encode(user)

}

func DelUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/user", getUsers).Methods("GET")
	r.HandleFunc("/api/user/{id}", getUsersbyId).Methods("GET")
	r.HandleFunc("/api/user", PostUsers).Methods("POST")
	r.HandleFunc("/api/user/{id}", PutUsers).Methods("PUT")
	r.HandleFunc("/api/user/{id}", DelUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r))
}
