package main

import (
	"RPN/api"

	"github.com/gorilla/mux"
)

// RegisterService creates mux router
func RegisterService() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/parse", api.RpnHandler).Methods("POST")
	router.HandleFunc("/health", api.HealthHandler).Methods("GET")
	return router
}
