package main

import (
	"encoding/json"
	"log"
	"net/http"

	"labix.org/v2/mgo/bson"

	"github.com/gorilla/mux"
)

var r *mux.Router

func init() {
	r = mux.NewRouter()
	r.HandleFunc("/contacts", getContactsHandler).Methods("GET")
	r.HandleFunc("/contacts/{id}", getContactHandler).Methods("GET")
}

// getContactsHandler returns a json response containing all
// the contacts.
//
// URL: GET /contacts
//
// For example, If there were only two contacts in total,
// GET /contacts and
// would yield the following response.
//
// Response:
// [
//  {
//    "id": "53a14760e3bdea4286000003",
//    "company": "Encom Inc.",
//    "person": "Sam Flynni !",
//    "email": "samflynn@encom.com",
//    "phone": "103-345-456",
//    "skypeid": "sam_flynn",
//    "country": "USA"
//    "slug": "sam-flynn"
//  },
//  {
//    "id": "53a14760e3bdea4286000004",
//    "company": "Encom Inc.",
//    "person": "Kevin Flynn",
//    "email": "kevinflynn@encom.com",
//    "phone": "234-877-988",
//    "skypeid": "kevin_flynn",
//    "country": "USA"
//  }
// ]
func getContactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c, err := GetAllContacts()
	if err != nil {
		log.Fatalf("%s", err)
	}
	enc := json.NewEncoder(w)
	if err = enc.Encode(c); err != nil {
		log.Println(err)
	}
}

// getContactHandler returns json data pertaining to a specific contact.
//
// URL: GET /contacts/{id}
//
// For example,
// GET /contacts/53a15a07e3bdea53d0000002
// Response:
// {
//   "id": "53a15a07e3bdea53d0000002",
//   "company": "Encom Inc.",
//   "person": "Flynn",
//   "email": "flynn@encom.com",
//   "country": "USA"
// }
func getContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	i := vars["id"]
	id := bson.ObjectIdHex(i)
	c, err := GetContact(id)
	if err != nil {
		log.Fatalf("%s", err)
	}
	enc := json.NewEncoder(w)
	if err = enc.Encode(c); err != nil {
		log.Println(err)
	}
}
