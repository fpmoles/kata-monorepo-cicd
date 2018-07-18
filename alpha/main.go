package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetAlpha).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetAlpha(w http.ResponseWriter, r *http.Request){
	message:= Message{Greeting: "Hello Alpha"}
	json.NewEncoder(w).Encode(message)
}

type Message struct{
	Greeting  string  `json:greeting,omitempty`
}
