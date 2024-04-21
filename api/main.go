package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kybouw/luhn/luhn"
)

func validate(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	number := vars["number"]
	json.NewEncoder(w).Encode(luhn.Validate(number))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/validate/{number}", validate).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
