package config

import (
	"log"
	"time"

	"labix.org/v2/mgo"
)

var (
	// Session is the configured MongoDB Session
	Session *mgo.Session

	// Db is the configured database
	Db *mgo.Database

	// LeadsCollection is the collection storing leads
	LeadsCollection *mgo.Collection

	// ContactsCollection is the collection storing contacts
	ContactsCollection *mgo.Collection

	// ContactsCollectionName is the name of the collection containing contacts.
	// It is required for forming the DB Query while performing text search.
	ContactsCollectionName string
)

// Configure configures the options, sets up the database,
// initializes exported session, database and collection variables.
func Configure(options map[string]string) {
	if options["name"] == "" {
		log.Fatalf("Configure requires the name of the MongoDB server")
	}
	if options["url"] == "" {
		log.Fatalf("Configure requires the url of the MongoDB server")
	}
	if options["leads"] == "" {
		log.Fatalf("Configure requires the name of the leads collection")
	}
	if options["contacts"] == "" {
		log.Fatalf("Configure requires the name of the contacts collection")
	}

	maxWait := time.Duration(5 * time.Second)
	session, err := mgo.DialWithTimeout(options["url"], maxWait)
	if err != nil {
		panic(err)
	}

	// In the Monotonic consistency mode reads may not be entirely up-to-date,
	// but they will always see the history of changes moving forward,
	// the data read will be consistent across sequential queries in the same session,
	// and modifications made within the session will be observed in following queries.
	session.SetMode(mgo.Monotonic, true)

	// Make the session check for errors, without imposing further constraints.
	session.SetSafe(&mgo.Safe{})

	Session = session
	Db = session.DB(options["name"])
	LeadsCollection = Db.C(options["leads"])
	ContactsCollection = Db.C(options["contacts"])
	ContactsCollectionName = options["contacts"]
	index := mgo.Index{
		Key: []string{"$text:person"},
	}
	err = ContactsCollection.EnsureIndex(index)
	if err != nil {
		log.Fatalln(err)
	}
}
