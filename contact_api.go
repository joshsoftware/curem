package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"labix.org/v2/mgo/bson"
)

var r *mux.Router

const ContactsBaseURL string = "http://localhost:3000/contacts/"

func init() {
	r = mux.NewRouter()
	r.HandleFunc("/contacts", getContactsHandler).Methods("GET")
	r.HandleFunc("/contacts", postContactsHandler).Methods("POST")
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
//    "person": "Sam Flynn",
//    "email": "samflynn@encom.com",
//    "phone": "103-345-456",
//    "skypeid": "sam_flynn",
//    "country": "USA"
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
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	enc := json.NewEncoder(w)
	if err = enc.Encode(c); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// postContactsHandler creates a new contact in the database.
//
// URL: POST /contacts
//
// Request:
// {
//   "company": "Encom Inc.",
//   "person": "Sam Flynn",
//   "email": "samflynn@encom.com",
//   "phone": "103-345-456",
//   "skypeid": "sam_flynn",
//   "country": "USA"
// }
//
// Response:
// HTTP/1.1 201 Created
// Location: http://localhost:3000/contacts/sam-flynn
func postContactsHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var c contact
	err := decoder.Decode(&c)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	n, err := NewContact(c.Company, c.Person, c.Email, c.Phone, c.SkypeId, c.Country)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	url := ContactsBaseURL + n.Slug
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusCreated)
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
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	enc := json.NewEncoder(w)
	if err = enc.Encode(c); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
