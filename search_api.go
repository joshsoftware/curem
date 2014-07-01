package main

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/joshsoftware/curem/config"
	"labix.org/v2/mgo/bson"
)

func init() {
	r.HandleFunc("/search", searchHandler).Methods("GET")
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()["q"][0]
	var c map[string]interface{}
	dbquery := bson.D{
		{"text", config.ContactsCollectionName},
		{"search", q},
	}
	err := config.Db.Run(dbquery, &c)
	if err != nil {
		log.Fatalf("%s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err = enc.Encode(c["results"]); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
