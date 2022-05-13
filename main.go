package main

import (
	"log"
	"net/http"
	"rest-go/serve"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	apirouter := r.PathPrefix("/api").Subrouter()
	apirouter.HandleFunc("/user", serve.GetAll).Methods("GET")
	apirouter.HandleFunc("/user/{id}", serve.GetOneUser).Methods("GET")
	apirouter.HandleFunc("/user", serve.PostUsers).Methods("POST")
	apirouter.HandleFunc("/update/{id}", serve.PutUsers).Methods("PUT")
	apirouter.HandleFunc("/delete/{id}", serve.DelUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
