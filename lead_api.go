package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"labix.org/v2/mgo/bson"
)

const LeadsBaseURL string = "http://localhost:3000/leads/"

func init() {
	r.HandleFunc("/leads", getLeadsHandler).Methods("GET")
	r.HandleFunc("/leads", postLeadHandler).Methods("POST")
	r.HandleFunc("/leads/{id}", getLeadHandler).Methods("GET")
	r.HandleFunc("/leads/{id}", patchLeadHandler).Methods("PATCH")
}

func getLeadsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c, err := GetAllLeads()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	if err = enc.Encode(c); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postLeadHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var l lead
	err := decoder.Decode(&l)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	n, err := NewLead(l.ContactId, l.Source, l.Owner, l.Status, l.TeamSize, l.RatePerHour,
		l.DurationInMonths, l.EstimatedStartDate, l.Comments)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	url := LeadsBaseURL + n.Id.Hex()
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusCreated)
}

func getLeadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	i := vars["id"]
	id := bson.ObjectIdHex(i)
	c, err := GetLead(id)
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

func patchLeadHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var l incomingLead
	err := decoder.Decode(&l)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	fl, err := GetLead(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = fl.copyIncomingFields(&l)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = fl.Update()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
