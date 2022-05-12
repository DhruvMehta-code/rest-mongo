package main

import (
	"log"
	"net/http"
	"rest-go/serve"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/user", serve.GetAll).Methods("GET")
	r.HandleFunc("/api/user/{id}", serve.GetOneUser).Methods("GET")
	r.HandleFunc("/api/user", serve.PostUsers).Methods("POST")
	r.HandleFunc("/api/user/{id}", serve.PutUsers).Methods("PUT")
	r.HandleFunc("/api/user/{id}", serve.DelUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
