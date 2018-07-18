package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetGamma).Methods("GET")
	log.Fatal(http.ListenAndServe(":8002", router))
}

func GetGamma(w http.ResponseWriter, r *http.Request){
	message:= Message{Greeting: "Hello Gamma"}
	json.NewEncoder(w).Encode(message)
}

type Message struct{
	Greeting  string  `json:greeting,omitempty`
}