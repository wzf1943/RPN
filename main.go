package main

import (
	"log"
	"net/http"

	"RPN/api"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/parse", api.RpnHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
