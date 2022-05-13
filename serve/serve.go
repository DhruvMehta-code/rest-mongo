package serve

import (
	"context"
	"encoding/json"

	"log"
	"net/http"
	"rest-go/helper"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = helper.ConnectDb()

func GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "apllication/json")

	var users []Users
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.GetError(err, w)
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var user Users

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

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "apllication/json")

	var user Users
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{
		"_id": id,
	}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user Users
	user.CreatedAt = time.Now()

	_ = json.NewDecoder(r.Body).Decode(&user)
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&user)
	if err != nil {
		helper.GetError(err, w)
	}
	json.NewEncoder(w).Encode(user)

}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user Users

	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&user)

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: user.Name},
		{Key: "email", Value: user.Email},
		{Key: "phone", Value: user.Phone},
		{Key: "description", Value: user.Desription},
		{Key: "CreatedAt", Value: time.Now()},
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

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)

}
