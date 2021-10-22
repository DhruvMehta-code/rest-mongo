package helper

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ErrorRes struct {
	Stat int    `json:"status"`
	Errm string `json:"message"`
}

func ConnectDb() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb+srv://unknown:12345@cluster0.b7h1w.mongodb.net/rest-go?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("rest-go").Collection("api")

	return collection
}

func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorRes{
		Errm: err.Error(),
		Stat: http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.Stat)
	w.Write(message)

}
