package main

import (
"encoding/json"
"github.com/gorilla/mux"
"log"
"net/http"
)

func GetHelloEndpoint(w http.ResponseWriter, r *http.Request) {
	params := "Hello world"
	json.NewEncoder(w).Encode(params)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", GetHelloEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}