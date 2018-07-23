package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetDelta).Methods("GET")
	log.Fatal(http.ListenAndServe(":8001", router))
}

func GetDelta(w http.ResponseWriter, r *http.Request){
	message:= Message{Greeting: "Hello Delta"}
	json.NewEncoder(w).Encode(message)
}

type Message struct{
	Greeting  string  `json:greeting,omitempty`
}