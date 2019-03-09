package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("hello I am go and this is an api, go luck.")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetPeopleEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":1528", router))
}



