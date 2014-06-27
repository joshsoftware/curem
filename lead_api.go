package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func init() {
	r.HandleFunc("/leads", getLeadsHandler).Methods("GET")
}

func getLeadsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c, err := GetAllLeads()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	enc := json.NewEncoder(w)
	if err = enc.Encode(c); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
