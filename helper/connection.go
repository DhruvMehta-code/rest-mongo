package helper

import (
	"context"
	"encoding/json"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ErrorRes struct {
	Stat int    `json:"status"`
	Errm string `json:"message"`
}

func ConnectDb() *mongo.Collection {

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	} 

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
        log.Fatal(err)
    }
	
	collection := client.Database("local").Collection("api")
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

